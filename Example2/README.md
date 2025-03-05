> ออกแบบ API ที่สามารถรองรับ 100,000 requests per second (RPS) ได้โดยไม่ให้เซิร์ฟเวอร์ล่ม และคำนึงถึงค่า server ด้วยเป็นหลัก โดยใช้ ใช้ techstack อะไรก็ได้
--- 1 request ต้อง ตรวจสอบ check user และ ตรวจสอบยอดเงิน จาก api ภายนอก

![flow](https://github.com/nice159123/example-test/blob/main/Example2/example2-flow.png)
![diagram](https://github.com/nice159123/example-test/blob/main/Example2/example2-diagram.png)
# Tech Stack ที่เลือก

| หัวข้อ        | เทคโนโลยี             | งาน
|---------------|-----------------------|----------
| Backend API   | Golang                | ดึงรายการจาก Queue มาจัดการ
| API Gateway   | Nginx                 | จัดการ Request ที่ส่งเข้ามา
| Queue         | RabbitMQ              | จัดการเก็บ Queue Request ที่เข้ามา
| Load Balancer | Nginx Load Balancer   | กระจาย Queue ไปยัง Backend
| Cache         | Redis                 | Caching External API Response
| Database      | PGpool-II, PostgreSQL | จัดการ Connection pooling