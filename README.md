# 🛠 BDS Square Backend
### Cài đặt Go trong WSL (chạy 1 lần duy nhất)

```bash
sudo snap install go --classic
```

---

### Build và khởi động toàn bộ dịch vụ (app + MySQL + Redis)

```bash
docker-compose up -d --build
```

---

### Chạy ứng dụng với hot reload (sử dụng air)

```bash
make run
```

---

### Debug ứng dụng Go (attach Delve)

```bash
make debug
```

---

### Debug bằng VS Code

1. Đảm bảo  đã cài **Go Extension by Google**
2. Vào tab Run & Debug (`Ctrl + Shift + D`)
3. Chọn cấu hình **Run Debug**
4. Nhấn `F5` để attach
5. Đặt breakpoint, gửi request đến API → debugger sẽ dừng đúng dòng

---

### Dừng debugger Delve (app vẫn tiếp tục chạy)

```bash
make stop-debug
```

---

### Generate SQL Queries (sử dụng sqlc)

Có thể generate các query Go từ file SQL thông qua:

```bash
sqlc generate
# hoặc
make sqlgen
```

---

### Cách tạo một table mới bằng goose:

```bash
goose -dir sql/schema create name_table sql
```

---

### Migrate database:

```bash
make up      # Thực thi tất cả các migration mới
make down    # Rollback migration gần nhất
make reset   # Xoá hết & migrate lại từ đầu
```

---

### Khởi tạo một Dependency Injection (DI):
1. Tạo một file name_controller.wire.go trong internal/wire và setup những DI cần dùng
2. **cd internal/wire** rồi chạy lệnh **wire**
3. Gọi wire vược được tạo tự động (wire_gen.go) ra trong router để gọi các func trong controller

### Swagger:
1. Link: http://localhost:8004/swagger/index.html
2. Mỗi lần viết swagger xong thì chạy **make swag**