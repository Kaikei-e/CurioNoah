import type React from "react";
import {
  Box,
  Button,
  Flex,
  Field,
  Input,
  Text,
} from "@chakra-ui/react";          // ← FormControl は削除
import { useNavigate } from "react-router-dom";
import { useAuth } from "../Auth/RequireAuth";
import type { User } from "../../lib/models/user";

export const Login = () => {
  const navigate = useNavigate();
  const auth = useAuth();

  function handleLogin(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const username = formData.get("username") as string;
    const password = formData.get("password") as string;

    if (!username || !password) {
      alert("Username or password cannot be empty");
      return;
    }
    const user: User = { username, password };
    if (auth) {
      auth.signin(user, () => navigate("/home"));
    } else {
      navigate("/");
    }
  }

  return (
    <Flex
      alignItems="center"
      justify="center"
      bgColor="#EAF2F8"
      w="100vw"
      h="100vh"
    >
      <Box
        maxW="320px"
        w="full"
        boxShadow="2xl"
        rounded="lg"
        p={6}
        textAlign="center"
        bgColor="#fff"
      >
        <Text fontSize="2xl" fontWeight="bold">
          Login
        </Text>
        <Box mt={4}>
          <form onSubmit={handleLogin}>
            {/* ユーザー名 */}
            <Field.Root required>
              <Field.Label>Username</Field.Label>
              <Input
                id="username"
                name="username"
                placeholder="Enter your username"
                fontFamily="Jost"
              />
              {/* 必要に応じてエラーメッセージ */}
              {/* <Field.ErrorText>Username is required</Field.ErrorText> */}
            </Field.Root>

            {/* パスワード */}
            <Field.Root mt={4} required>
              <Field.Label>Password</Field.Label>
              <Input
                id="password"
                name="password"
                type="password"
                placeholder="Enter your password"
                fontFamily="Jost"
              />
            </Field.Root>

            {/* 送信ボタン */}
            <Button
              colorScheme="blue"
              mt={6}
              type="submit"
              width="full"
              fontFamily="Jost"
            >
              Login
            </Button>
          </form>
        </Box>
      </Box>
    </Flex>
  );
};
