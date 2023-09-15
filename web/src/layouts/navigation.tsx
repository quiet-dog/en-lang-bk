import NavCard from "@/components/NavCard";
import { NavCardProps } from "@/components/NavCard/types";
import { commonStyle } from "@/styles/common";
import { Col, Layout, Row, Space } from "antd";
import { Content, Footer, Header } from "antd/es/layout/layout";

export default function Navigation() {
    const cardProps: NavCardProps[] = [
        {
            imgUrl: "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
            title: "测试props",
            onClick: () => {
                console.log(1);
            },
        },
        {
            imgUrl: "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
            title: "测试props",
            onClick: () => {
                console.log(2);
            },
        },
        {
            imgUrl: "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
            title: "测试props",
            onClick: () => {
                console.log(3);
            },
        },
    ];

    const selectItems = cardProps.map((item, index) => (
        <Col span={Math.floor(24 / cardProps.length)}>
            <NavCard title={item.title} imgUrl={item.imgUrl} onClick={item.onClick} key={index} />
        </Col>
    ));

    return (
        <>
            <Space direction="vertical" style={{ width: "100%" }}>
                <Layout style={commonStyle}>
                    <Header style={commonStyle}></Header>
                    <Content style={commonStyle}>
                        <Row justify="space-around" gutter={16}>
                            {selectItems}
                        </Row>
                    </Content>
                    <Footer style={commonStyle}></Footer>
                </Layout>
            </Space>
        </>
    );
}
