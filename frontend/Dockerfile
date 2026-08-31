FROM node:20-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 5173

# Host 0.0.0.0 обов'язковий, щоб Vite приймав з'єднання звне контейнера
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]