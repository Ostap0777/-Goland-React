package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

const carSelect = `
SELECT c.id, m.name, c.model, c.year, c.price, c.mileage, c.description, c.seller_name, c.created_at
FROM cars c
JOIN makes m ON m.id = c.make_id`

type PostgresRepository struct {
   db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
   return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, car *Car) error {
   tx, err := r.db.BeginTx(ctx, nil)
   if err != nil {
      log.Printf("[ERROR Create Tx Begin]: %v", err)
      return err
   }
   defer tx.Rollback()

   carQuery := `
INSERT INTO cars (make_id, model, year, price, mileage, description, seller_name)
SELECT m.id, $2, $3, $4, $5, $6, $7
FROM makes m
WHERE LOWER(m.name) = LOWER($1)
RETURNING id, created_at`

   err = tx.QueryRowContext(
      ctx, carQuery,
      car.Make, car.Model, car.Year, car.Price, car.Mileage, car.Description, car.SellerName,
   ).Scan(&car.ID, &car.CreatedAt)

   if errors.Is(err, sql.ErrNoRows) {
      log.Printf("[ERROR Create Car]: Unknown make '%s'", car.Make)
      return ErrUnknownMake
   }
   if err != nil {
      log.Printf("[ERROR Create Car Query]: %v", err)
      return fmt.Errorf("failed to insert car: %w", err)
   }

   if len(car.Images) > 0 {
      imgQuery := `
INSERT INTO car_images (car_id, url, is_main, "order")
VALUES ($1, $2, $3, $4)
RETURNING id, created_at`

      for i := range car.Images {
         car.Images[i].CarID = car.ID
         err := tx.QueryRowContext(
            ctx, imgQuery,
            car.Images[i].CarID, car.Images[i].URL, car.Images[i].IsMain, car.Images[i].Order,
         ).Scan(&car.Images[i].ID, &car.Images[i].CreatedAt)
         if err != nil {
            log.Printf("[ERROR Create Image Query index %d]: %v", i, err)
            return fmt.Errorf("failed to insert image: %w", err)
         }
      }
   }

   if err := tx.Commit(); err != nil {
      log.Printf("[ERROR Create Tx Commit]: %v", err)
      return err
   }

   return nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int64) (*Car, error) {
   var c Car
   err := r.db.QueryRowContext(ctx, carSelect+` WHERE c.id = $1`, id).Scan(
      &c.ID, &c.Make, &c.Model, &c.Year, &c.Price, &c.Mileage, &c.Description, &c.SellerName, &c.CreatedAt,
   )
   if errors.Is(err, sql.ErrNoRows) {
      return nil, ErrNotFound
   }
   if err != nil {
      return nil, err
   }

   images, err := r.getImagesByCarID(ctx, c.ID)
   if err != nil {
      return nil, err
   }
   c.Images = images

   return &c, nil
}

func (r *PostgresRepository) List(ctx context.Context, filter Filter) ([]Car, error) {
   var (
      b    strings.Builder
      args []any
   )
   b.WriteString(carSelect)
   b.WriteString(" WHERE 1=1")

   add := func(sqlExpr string, value any) {
      args = append(args, value)
      fmt.Fprintf(&b, " AND "+sqlExpr, len(args))
   }

   if filter.Make != "" {
      add("LOWER(m.name) = LOWER($%d)", filter.Make)
   }
   if filter.Model != "" {
      add("LOWER(c.model) = LOWER($%d)", filter.Model)
   }
   if filter.MinYear > 0 {
      add("c.year >= $%d", filter.MinYear)
   }
   if filter.MaxYear > 0 {
      add("c.year <= $%d", filter.MaxYear)
   }
   if filter.MinPrice > 0 {
      add("c.price >= $%d", filter.MinPrice)
   }
   if filter.MaxPrice > 0 {
      add("c.price <= $%d", filter.MaxPrice)
   }

   args = append(args, filter.Limit, filter.Offset)
   fmt.Fprintf(&b, " ORDER BY c.created_at DESC LIMIT $%d OFFSET $%d", len(args)-1, len(args))

   rows, err := r.db.QueryContext(ctx, b.String(), args...)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   out := make([]Car, 0)
   for rows.Next() {
      var c Car
      if err := rows.Scan(
         &c.ID, &c.Make, &c.Model, &c.Year, &c.Price, &c.Mileage, &c.Description, &c.SellerName, &c.CreatedAt,
      ); err != nil {
         return nil, err
      }
      out = append(out, c)
   }
   if err := rows.Err(); err != nil {
      return nil, err
   }

   for i := range out {
      images, err := r.getImagesByCarID(ctx, out[i].ID)
      if err != nil {
         return nil, err
      }
      out[i].Images = images
   }

   return out, nil
}

func (r *PostgresRepository) Update(ctx context.Context, car *Car) error {
   tx, err := r.db.BeginTx(ctx, nil)
   if err != nil {
      return err
   }
   defer tx.Rollback()

   query := `
UPDATE cars c
SET make_id = m.id,
    model = $3,
    year = $4,
    price = $5,
    mileage = $6,
    description = $7,
    seller_name = $8
FROM makes m
WHERE c.id = $1 AND LOWER(m.name) = LOWER($2)
RETURNING c.id, m.name, c.model, c.year, c.price, c.mileage, c.description, c.seller_name, c.created_at`

   err = tx.QueryRowContext(
      ctx, query,
      car.ID, car.Make, car.Model, car.Year, car.Price, car.Mileage, car.Description, car.SellerName,
   ).Scan(
      &car.ID, &car.Make, &car.Model, &car.Year, &car.Price, &car.Mileage, &car.Description, &car.SellerName, &car.CreatedAt,
   )
   if errors.Is(err, sql.ErrNoRows) {
      exists, lookupErr := r.exists(ctx, car.ID)
      if lookupErr != nil {
         return lookupErr
      }
      if !exists {
         return ErrNotFound
      }
      return ErrUnknownMake
   }
   if err != nil {
      return err
   }

   if len(car.Images) > 0 {
      if _, err := tx.ExecContext(ctx, `DELETE FROM car_images WHERE car_id = $1`, car.ID); err != nil {
         return err
      }

      imgQuery := `
INSERT INTO car_images (car_id, url, is_main, "order")
VALUES ($1, $2, $3, $4)
RETURNING id, created_at`

      for i := range car.Images {
         car.Images[i].CarID = car.ID
         err := tx.QueryRowContext(
            ctx, imgQuery,
            car.Images[i].CarID, car.Images[i].URL, car.Images[i].IsMain, car.Images[i].Order,
         ).Scan(&car.Images[i].ID, &car.Images[i].CreatedAt)
         if err != nil {
            return err
         }
      }
   }

   return tx.Commit()
}

func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
   res, err := r.db.ExecContext(ctx, `DELETE FROM cars WHERE id = $1`, id)
   if err != nil {
      return err
   }
   n, err := res.RowsAffected()
   if err != nil {
      return err
   }
   if n == 0 {
      return ErrNotFound
   }
   return nil
}

func (r *PostgresRepository) exists(ctx context.Context, id int64) (bool, error) {
   var found bool
   err := r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM cars WHERE id = $1)`, id).Scan(&found)
   return found, err
}

func (r *PostgresRepository) getImagesByCarID(ctx context.Context, carID int64) ([]CarImage, error) {
   query := `
SELECT id, car_id, url, is_main, "order", created_at
FROM car_images
WHERE car_id = $1
ORDER BY "order" ASC, id ASC`

   rows, err := r.db.QueryContext(ctx, query, carID)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   images := make([]CarImage, 0)
   for rows.Next() {
      var img CarImage
      if err := rows.Scan(&img.ID, &img.CarID, &img.URL, &img.IsMain, &img.Order, &img.CreatedAt); err != nil {
         return nil, err
      }
      images = append(images, img)
   }
   return images, rows.Err()
}