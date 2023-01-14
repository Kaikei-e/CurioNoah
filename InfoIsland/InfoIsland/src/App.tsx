import './App.css'
// /images/explore_with_ship.jpg
import {Stack, Text, Box, Button, Spacer, Flex} from '@chakra-ui/react'
import {ArrowForwardIcon} from '@chakra-ui/icons'

export const App = () => (
    <Box w={'100vw'} h={'100vh'}
         minW={'100%'} minH={'100%'}
         backgroundImage={"/images/ship-sunset.jpg"}
         backgroundSize={'cover'}
         backgroundRepeat={'no-repeat'}
         p={10}
    >
        <Flex minWidth='max-content' h={'100%'} w={'100%'}>
            <Stack
                direction="column"
                justify="flex-start"
                align="flex-start"
                spacing="10px"
            >
                <Text
                    fontFamily="Jost"
                    fontWeight="regular"
                    fontSize={{base: "3xl", md: "4xl", lg: "80px"}}
                    color="#FFF"
                    width={{base: "100%", md: "100%", lg: "100%"}}
                    maxWidth="100%"
                >
                    <span>CurioNoah </span>
                    <br/>
                    <Box as="span" fontWeight="light"
                         fontSize={{base: "xl", md: "2xl", lg: "6xl"}}>
                        Explore, Store, Extend
                    </Box>
                    <br/>
                    <Box as={"span"} fontWeight="light" fontSize={{base: "xl", md: "2xl", lg: "6xl"}}>
                        : by curiosity.
                    </Box>
                </Text>
            </Stack>
            <Spacer/>
            <Flex flexDir={'column'} justifyContent={'flex-end'} p='4%'>
                <Button size={{base: "lg", sm: "sm", md: "md", lg: "lg"}}
                        variant='outline'
                        rightIcon={<ArrowForwardIcon
                            color={'#FFF'}
                            boxSize={{base: "xl", md: "2xl", lg: "20px"}}
                        />}
                >

                    <Text textColor={'#FFF'}
                          fontFamily="Jost"
                          fontSize={{base: "md", md: "2md", lg: "2xl"}}>
                        Explore
                    </Text>
                </Button>
            </Flex>
        </Flex>
    </Box>

)

export default App;