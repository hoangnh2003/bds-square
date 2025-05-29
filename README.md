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

1. Đảm bảo bạn đã cài **Go Extension by Google**
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