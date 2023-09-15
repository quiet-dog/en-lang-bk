import { ConfigProvider } from "antd";

function App() {
	return (
		<>
			<ConfigProvider>
				<Outlet />
			</ConfigProvider>
		</>
	);
}

export default App;
