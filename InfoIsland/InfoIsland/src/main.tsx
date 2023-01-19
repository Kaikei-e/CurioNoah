import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import {ChakraProvider} from "@chakra-ui/react";
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import Home from "./home/Home";
import InsightStreamMain from "./components/InshitStream/InsightStreamMain";


const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
    },
    {
        path: "/home",
        element: <Home/>,
    },
    {
        path: "/insight-stream",
        element: <InsightStreamMain/>,
    }

]);

const rootElement = document.getElementById('root')
// @ts-ignore
ReactDOM.createRoot(rootElement).render(
    <React.StrictMode>
        <ChakraProvider>
            <RouterProvider router={router}/>
        </ChakraProvider>
    </React.StrictMode>,
)


// https://coolors.co/eaf2f8-a9d0f5-85c1e9-71aedc-5d8dcf