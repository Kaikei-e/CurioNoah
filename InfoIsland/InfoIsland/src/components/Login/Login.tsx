import React, { useState } from "react";
import {
  Box,
  Button,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Text,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../Auth/RequireAuth";
import { User } from "../../lib/models/user";

export const Login = () => {
  const navigate = useNavigate();
  const auth = useAuth();

  function handleLogin(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    let formData = new FormData(event.currentTarget);
    let username = formData.get("username") as string;
    let password = formData.get("password") as string;

    if (username === "" || password === "") {
      alert("Username or password cannot be empty");
      return;
    }

    const user: User = { username: username, password: password };

    if (auth !== null) {
      auth.signin(user, () => {
        navigate("/home");
      });
    } else {
      navigate("/");
    }
  }

  return (
    <Flex
      alignItems={"center"}
      justify={"center"}
      bgColor={"#EAF2F8"}
      w={"100vw"}
      h={"100vh"}
      minH={"100%"}
      minW={"100%"}
    >
      <Box
        maxW={"320px"}
        w={"full"}
        boxShadow={"2xl"}
        rounded={"lg"}
        p={6}
        textAlign={"center"}
        bgColor={"#ffffff"}
      >
        <Text fontSize={"2xl"} fontWeight={"bold"}>
          Login
        </Text>
        <Box mt={4}>
          <form onSubmit={handleLogin}>
            <FormControl>
              <FormLabel fontFamily={"Jost"}>Username</FormLabel>
              <Input
                id={"username"}
                name={"username"}
                type="text"
                placeholder="Enter your username"
                fontFamily={"Jost"}
              />
              <FormLabel fontFamily={"Jost"}>Password</FormLabel>
              <Input
                id={"password"}
                name={"password"}
                type="password"
                placeholder="Enter your password"
                fontFamily={"Jost"}
              />
              <Button
                colorScheme="blue"
                mt={6}
                type="submit"
                width={"full"}
                fontFamily={"Jost"}
              >
                Login
              </Button>
            </FormControl>
          </form>
        </Box>
      </Box>
    </Flex>
  );
};

