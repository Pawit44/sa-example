import { lazy } from "react";


import type { RouteObject } from "react-router-dom";

import MinimalLayout from "../layout/MinimalLayout";

import Loadable from "../components/third-patry/Loadable";


const MainPages = Loadable(lazy(() => import("../pages/authentication/Login")));

const Registerages = Loadable(

  lazy(() => import("../pages/authentication/Register"))

);


const MainRoutes = (): RouteObject => {

  return {

    path: "/",

    element: <MinimalLayout />,

    children: [

      {

        path: "/",

        element: <MainPages />,

      },

      {

        path: "/signup",

        element: <Registerages />,

      },

      {

        path: "*", // เป็นตัวที่บอกว่าถ้าพิมพ์อย่างอื่นที่ไม่ใช่ / และ /signup ทั้งหมดจะให้ link ไปหน้า login ทั้งหมด

        element: <MainPages />,

      },

    ],

  };

};


export default MainRoutes;