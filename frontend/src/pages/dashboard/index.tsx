import { Col, Row, Card, Statistic, Table } from "antd";
import { useState , useEffect } from "react";
import axios from "axios";

import {

  AuditOutlined,

  UserOutlined,

  PieChartOutlined,

  StockOutlined,

} from "@ant-design/icons";

import type { ColumnsType } from "antd/es/table";


interface DataType {

  key: string,
  ID: string,
  first_name: string,
  last_name: string,
  email: string,
  age: number,
  address: string,

}


const columns: ColumnsType<DataType> = [

  {

    title: "ลำดับ",

    dataIndex: "ID",

    key: "id",

  },

  {

    title: "ชื่อ",

    dataIndex: "first_name",

    key: "firstname",

  },

  {

    title: "นามสกุล",

    dataIndex: "last_name",

    key: "lastname",

  },

  {

    title: "อีเมล",

    dataIndex: "email",

    key: "email",

  },

  {

    title: "อายุ",

    dataIndex: "age",

    key: "age",

  },
  {

    title: "ที่อยู่",

    dataIndex: "address",

    key: "address",

  },

];



export default function index() {
  const [data,setData] = useState<DataType []>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {

    setLoading(true)
    setError(null)
    
    axios.get("http://localhost:8000/api/users/latest")
      .then((res) => {
        const withKey = res.data.map((item: any, index: number) => ({
          ...item,
          key: item.ID || index.toString(),
        }));
        setData(withKey);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Error fetching data:", err);
        setError("ไม่สามารถโหลดข้อมูลได้");
        setLoading(false);
      });
  }, []);


  return (

    <>

      <Row gutter={[16, 16]}>

        <Col xs={24} sm={24} md={24} lg={24} xl={24}>

          <h2>แดชบอร์ด</h2>

        </Col>


        <Col xs={24} sm={24} md={24} lg={24} xl={24}>

          <Card style={{ backgroundColor: "#f5f5f5" }}>

            <Row gutter={[16,16]}>

              <Col xs={24} sm={24} md={12} lg={12} xl={6}>

                <Card

                  bordered={false}

                  style={{

                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",

                  }}

                >

                  <Statistic

                    title="จำนวน"

                    value={1800}

                    prefix={<StockOutlined />}

                  />

                </Card>

              </Col>


              <Col xs={24} sm={24} md={12} lg={12} xl={6}>

                <Card

                  bordered={false}

                  style={{

                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",

                  }}

                >

                  <Statistic

                    title="จำนวน"

                    value={200}

                    valueStyle={{ color: "back" }}

                    prefix={<AuditOutlined />}

                  />

                </Card>

              </Col>


              <Col xs={24} sm={24} md={12} lg={12} xl={6}>

                <Card

                  bordered={false}

                  style={{

                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",

                  }}

                >

                  <Statistic

                    title="จำนวน"

                    value={3000}

                    valueStyle={{ color: "black" }}

                    prefix={<PieChartOutlined />}

                  />

                </Card>

              </Col>


              <Col xs={24} sm={24} md={12} lg={12} xl={6}>

                <Card

                  bordered={false}

                  style={{

                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",

                  }}

                >

                  <Statistic

                    title="จำนวน"

                    value={10}

                    valueStyle={{ color: "black" }}

                    prefix={<UserOutlined />}

                  />

                </Card>

              </Col>

            </Row>

          </Card>

        </Col>


        <Col xs={24} sm={24} md={24} lg={24} xl={24}>

          <h3>ผู้ใช้งานล่าสุด</h3>

        </Col>


        <Col xs={24} sm={24} md={24} lg={24} xl={24}>

          {error && (
            <div style={{ color: 'red', textAlign: 'center', padding: '20px' }}>
              {error}
            </div>
          )}
          
          <Table 
            columns={columns} 
            dataSource={data} 
            loading={loading}
            locale={{
              emptyText: loading ? 'กำลังโหลด...' : 'ไม่มีข้อมูล'
            }}
          />

        </Col>

      </Row>

    </>

  );

}