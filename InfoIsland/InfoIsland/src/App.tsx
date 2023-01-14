import './App.css'
// /images/explore_with_ship.jpg
import {Stack, Text, Box} from '@chakra-ui/react'

export const App = () => (
    <Box w={'100vw'} h={'100vh'} minW={'100%'} minH={'100%'} backgroundImage={"/images/explore_with_ship.jpg"}>

        <Stack
            direction="column"
            justify="flex-start"
            align="flex-start"
            spacing="10px"
        >
            <Box>
                <Text
                    fontFamily="Jost"
                    fontWeight="regular"
                    fontSize="97.2px"
                    color="#000000"
                    width="1174px"
                    maxWidth="100%"
                >
                    <span>CurioNoah </span>
                    <Box as="span" fontWeight="light" fontSize="81px">
                        Explore, Store, Extend: by curiosity.
                    </Box>
                </Text>
            </Box>
        </Stack>
    </Box>

)

export default App;