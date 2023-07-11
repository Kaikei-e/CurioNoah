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

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const handleLogin = (event: { preventDefault: () => void }) => {
    event.preventDefault();
    // perform login logic here

    if (username === "" || password === "") {
      alert("Username or password cannot be empty");
      return;
    }

    if (username === "admin" || password === "admin") {
      navigate("/home");
      return;
    }
    console.log("Username: ", username);
    console.log("Password: ", password);
  };

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
              <FormLabel>Username</FormLabel>
              <Input
                type="text"
                placeholder="Enter your username"
                onChange={(event) => setUsername(event.target.value)}
              />
            </FormControl>
            <FormControl mt={6}>
              <FormLabel>Password</FormLabel>
              <Input
                type="password"
                placeholder="Enter your password"
                onChange={(event) => setPassword(event.target.value)}
              />
            </FormControl>
            <Button colorScheme="blue" mt={6} type="submit" width={"full"}
                onSubmit={handleLogin}
            >
              Login
            </Button>
          </form>
        </Box>
      </Box>
    </Flex>
  );
};

export default Login;
