import { ConfigProvider } from "antd";
import { NextUIProvider } from "@nextui-org/system";
import { Button } from "@nextui-org/button";

function App() {
  return (
    <>
      <NextUIProvider>
        <Button>Press me</Button>

        <ConfigProvider>
          <Outlet />
        </ConfigProvider>
      </NextUIProvider>
    </>
  );
}

export default App;
