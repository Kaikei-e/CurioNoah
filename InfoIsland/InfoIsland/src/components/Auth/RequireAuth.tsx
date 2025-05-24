import React, { useState, createContext, useContext, useEffect } from "react";
import { Navigate, useLocation, useNavigate } from "react-router-dom";
import { User } from "../../lib/models/user";

interface AuthContextType {
  user: User | null;
  signin: (user: User, callback: () => void) => void;
  signout: (callback: () => void) => void;
  isAuthenticating: boolean;
}
// Use `User | null` as the context type and initialize it to null.
let AuthContext = createContext<AuthContextType | null>(null);

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}

export function RequireAuth({ children }: { children: React.ReactNode }) {
  const { user, signin, signout, isAuthenticating } = useAuth();
  let location = useLocation();
  const navigate = useNavigate();

  if (isAuthenticating) {
    // Here you could return a loading spinner or just null, depending on what you want.
    return null;
  }

  if (!user) {
    return <Navigate to="/" state={{ from: location }} replace />;
  }

  return <>{children}</>;
}

type Props = {
  children: React.ReactNode;
};

export const AuthProvider: React.VFC<Props> = (props) => {
  let [user, setUser] = useState<User | null>(null);
  let [isAuthenticating, setAuthenticating] = useState(true);

  useEffect(() => {
    const user = sessionStorage.getItem("user");
    if (user) {
      setUser(JSON.parse(user));
    }
    setAuthenticating(false); // Authentication check is finished.
  }, []);

  const signin = (user: User, callback: () => void) => {
    if (user.username === "admin" && user.password === "admin") {
      setUser({ username: "admin", password: "admin" });
      sessionStorage.setItem(
        "user",
        JSON.stringify({ username: "admin", password: "admin" }),
      );
      callback();
    }
  };

  const signout = (callback: () => void) => {
    setUser(null);
    sessionStorage.removeItem("user");
    callback();
  };

  let value = { user, signin, signout, isAuthenticating };

  return (
    <>
      <AuthContext.Provider value={value}>
        {props.children}
      </AuthContext.Provider>
    </>
  );
};
