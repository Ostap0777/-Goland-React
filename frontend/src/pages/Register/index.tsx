import React, { useEffect, useState } from "react";
import Input from "../../shared/UI/Input";
import Button from "../../shared/UI/Button";
import styles from "./styles.module.scss";

function RegisterPage() {
  const [login, setLogin] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  useEffect(() => {
    console.log(login);
    console.log(password);
  }, [login, password]);
  return (
    <div className={styles.wrapper}>
      <div className={styles.card}>
        <h2 className={styles.title}>Register</h2>
        <form className={styles.form}>
          <Input
            type="email"
            placeholder="Email or Phone"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
          />
          <Input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button title="Create" type="submit" />
        </form>
      </div>
    </div>
  );
}

export default RegisterPage;
