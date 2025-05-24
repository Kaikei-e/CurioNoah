import { Flex, Text, Button, Table } from "@chakra-ui/react";
import { AddFeed } from "./AddFeed";
import { InfiniteFeeds } from "./InsightStream/InfiniteFeeds";
import { collectAllNewFeeds } from "../../../lib/fetcher/collectAllNewFeeds";

export const InsightStreamMain = () => {
  return (
    <Flex
      h="100vh"
      w="100vw"
      flexDirection="row"
      fontFamily="Jost"
      boxSizing="border-box"
    >
      {/* サイドバー */}
      <Flex flexDirection="column" w="25%" h="100%" bgColor="#88A6B1">
        <Flex
          h="6%"
          justifyContent="center"
          alignItems="center"
          rounded="xl"
          p="4%"
          textAlign="center"
        >
          <Text
            borderBottom="2px solid #000"
            fontWeight="regular"
            fontSize={{ base: "lg", md: "xl", lg: "4xl" }}
            color="#fff"
            textAlign="start"
          >
            InsightStream
          </Text>
        </Flex>

        <Flex
          flexDirection="column"
          w="95%"
          h="94%"
          alignItems="center"
          p="4%"
        >
          <Flex
            flexDirection="column"
            w="80%"
            h="60%"
            bgColor="#EAF2F8"
            rounded="xl"
            border="1px solid #000"
          >
            <Text
              fontWeight="regular"
              fontSize={{ base: "xl", md: "xl", lg: "2xl" }}
              color="#000"
              m="2%"
              textAlign="center"
            >
              Status
            </Text>

            {/* テーブル: Chakra UI v3 実装 */}
            <Table.ScrollArea w="100%" h="auto">
              <Table.Root 
                variant="outline" 
                size="md"
              >
                <Table.Header>
                  <Table.Row>
                    <Table.ColumnHeader>Topic</Table.ColumnHeader>
                    <Table.ColumnHeader textAlign="end">Statistic</Table.ColumnHeader>
                  </Table.Row>
                </Table.Header>
                <Table.Body>
                  <Table.Row>
                    <Table.Cell>Following</Table.Cell>
                    <Table.Cell textAlign="end">20</Table.Cell>
                  </Table.Row>
                  <Table.Row>
                    <Table.Cell>
                      Number of feeds
                      <br />
                      read by now
                    </Table.Cell>
                    <Table.Cell textAlign="end">100</Table.Cell>
                  </Table.Row>
                </Table.Body>
              </Table.Root>
            </Table.ScrollArea>
          </Flex>
        </Flex>
      </Flex>

      {/* メインコンテンツ */}
      <Flex
        flexDirection="row"
        w="75%"
        h="100%"
        bgColor="#88A6B1"
        p="2%"
        justifyContent="space-evenly"
      >
        <Flex
          flexDirection="column"
          w="70%"
          h="100%"
          bgColor="#EAF2F8"
          rounded="xl"
          p="2%"
        >
          <Flex
            display="flex"
            w="100%"
            h="100%"
            overflowY="scroll"
            flexDirection="column"
          >
            <InfiniteFeeds />
          </Flex>
        </Flex>

        <Flex flexDirection="column" w="30%" h="100%" p="2%">
          <Flex
            w="100%"
            h="45%"
            p="2%"
            rounded="xl"
            border="1px solid #000"
            shadow="md"
          >
            <AddFeed />
          </Flex>
          <Flex
            w="100%"
            h="10%"
            rounded="xl"
            p="2%"
            mt="2%"
            border="1px solid #000"
            shadow="md"
            justifyContent="center"
            alignItems="center"
          >
            <form>
              <Button
                onClick={async () => {
                  await collectAllNewFeeds();
                  alert("Collecting new feeds!");
                }}
              >
                Collect New Feeds
              </Button>
            </form>
          </Flex>
        </Flex>
      </Flex>
    </Flex>
  );
};