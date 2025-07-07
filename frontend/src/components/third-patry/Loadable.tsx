import type { ComponentType } from "react";
import { Suspense } from "react";

import Loader from "./Loader";


const Loadable =

  <P extends object>(Component: ComponentType<P>): ComponentType<P> =>

  (props: P) =>

    (

      <Suspense fallback={<Loader />}>

        <Component {...props} />

      </Suspense>

    );


export default Loadable;