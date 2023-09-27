import { ConfigProvider } from "antd";
import { NextUIProvider } from "@nextui-org/system";

function App() {
  return (
    <>
      <NextUIProvider>
        <ConfigProvider>
          <Outlet />
        </ConfigProvider>
      </NextUIProvider>
    </>
  );
}

export default App;
