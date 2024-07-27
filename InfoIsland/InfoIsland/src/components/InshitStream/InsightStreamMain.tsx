import {
  Flex,
  Text,
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableCaption,
  TableContainer,
  Menu,
  MenuButton,
  IconButton,
  MenuList,
  MenuItem,
} from "@chakra-ui/react";
import { AddIcon, HamburgerIcon, EditIcon } from "@chakra-ui/icons";
import { AddFeed } from "../Islands/InsightStream/AddFeed";

export const InsightStreamMain = () => {
  return (
    <div>
      <Flex
        h={"100vh"}
        w={"100vw"}
        minH={"100%"}
        minW={"100%"}
        overflowX={"hidden"}
        flexDirection={"row"}
        fontFamily="Jost"
      >
        <Flex flexDirection={"column"} w={"20%"} h={"100%"} bgColor={"#88A6B1"}>
          <Flex
            h={"6%"}
            justifyContent={"center"}
            alignItems={"center"}
            rounded={"xl"}
            p={"4%"}
            textAlign={"center"}
          >
            <Text
              style={{
                borderBottom: "2px solid #000",
              }}
              fontWeight="regular"
              fontSize={{ base: "lg", md: "xl", lg: "4xl" }}
              color="#fff"
              textAlign={"start"}
            >
              InsightStream
            </Text>
          </Flex>
          <Flex
            flexDirection={"column"}
            w={"100%"}
            h={"94%"}
            alignItems={"center"}
            p={"4%"}
          >
            <Flex
              flexDirection={"column"}
              w={"80%"}
              h={"60%"}
              bgColor={"#EAF2F8"}
              rounded={"xl"}
              border={"1px solid #000"}
            >
              <Text
                fontWeight="regular"
                fontSize={{ base: "xl", md: "xl", lg: "2xl" }}
                color="#000"
                m={"2%"}
                textAlign={"center"}
              >
                Status
              </Text>
              <TableContainer w={"100%"}>
                <Table variant="simple">
                  <TableCaption>You can find the stats of reading</TableCaption>
                  <Thead>
                    <Tr>
                      <Th>Topic</Th>
                      <Th isNumeric>Statistic</Th>
                    </Tr>
                  </Thead>
                  <Tbody>
                    <Tr>
                      <Td>Following</Td>
                      <Td isNumeric>20</Td>
                    </Tr>
                    <Tr>
                      <Td>
                        Number of feeds <br />
                        read by now
                      </Td>
                      <Td isNumeric>100</Td>
                    </Tr>
                  </Tbody>
                </Table>
              </TableContainer>
            </Flex>

            <Flex
              flexDirection={"column"}
              h={"30%"}
              w={"80%"}
              mt={"4%"}
              justify-content={"flex-start"}
              alignItems={"flex-start"}
            >
              <Menu>
                <MenuButton
                  as={IconButton}
                  aria-label="Options"
                  icon={<HamburgerIcon />}
                  variant="outline"
                />
                <MenuList>
                  <MenuItem icon={<AddIcon />} command="⌘T">
                    New Tab
                  </MenuItem>
                  <MenuItem icon={<AddIcon />} command="⌘N">
                    New Window
                  </MenuItem>
                  <MenuItem icon={<AddIcon />} command="⌘⇧N">
                    Open Closed Tab
                  </MenuItem>
                  <MenuItem icon={<EditIcon />} command="⌘O">
                    Open File...
                  </MenuItem>
                </MenuList>
              </Menu>
            </Flex>
          </Flex>
        </Flex>
        <Flex
          flexDirection={"row"}
          w={"80%"}
          h={"100%"}
          bgColor={"#88A6B1"}
          p={"2%"}
        >
          <Flex
            flexDirection={"column"}
            w={"70%"}
            h={"100%"}
            bgColor={"#EAF2F8"}
            rounded={"xl"}
            p={"2%"}
          >
            {/*<StreamLine/>*/}
            <Flex>
              <p>StreamLine</p>
            </Flex>
          </Flex>
          <Flex
            w={"25%"}
            h={"40%"}
            rounded={"xl"}
            p={"2%"}
            ml={"5%"}
            border={"1px solid #000"}
            shadow={"md"}
          >
            <AddFeed />
          </Flex>
        </Flex>
      </Flex>
    </div>
  );
};
