import React, { useState, createContext, useContext } from "react";
import { Navigate, useLocation, useNavigate } from "react-router-dom";
import { User } from "../../lib/models/user";

interface AuthContextType {
  user: User | null;
  signin: (user: User, callback: () => void) => void;
  signout: (callback: () => void) => void;
}

// Use `User | null` as the context type and initialize it to null.
let AuthContext = createContext<AuthContextType | null>(null);

export function useAuth() {
  return useContext(AuthContext);
}

export function RequireAuth({
  children,
}: {
  children: React.ReactNode;
}) {
  let auth = useAuth();
  let location = useLocation();
  const navigate = useNavigate();

  if (auth === null) {
    navigate("/");
  } else if (!auth.user) {
    // Redirect them to the /login page, but save the current location they were
    // trying to go to when they were redirected. This allows us to send them
    // along to that page after they login, which is a nicer user experience
    // than dropping them off on the home page.
    return <Navigate to="/" state={{ from: location }} replace />;
  }

  return <> {children} </>;
}

type Props = {
  children: React.ReactNode;
};

export const AuthProvider: React.VFC<Props> = (props) => {
  let [user, setUser] = useState<User | null>(null);

  const signin = (user: User, callback: () => void) => {
    if (user.username === "admin" && user.password === "admin") {
      setUser({ username: "admin", password: "admin" });
      callback();
    }
  };

  const signout = (callback: () => void) => {
    // Insert your signout logic here.
    // If signout is successful, unset the user state.
    setUser(null);
    callback();
  };

  let value = { user, signin, signout };

  return (
    <>
      <AuthContext.Provider value={value}>
        {props.children}
      </AuthContext.Provider>
    </>
  );
};
