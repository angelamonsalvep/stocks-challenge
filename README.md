# 🧩 Stocks Challenge

Este proyecto implementa un sistema que obtiene información de acciones desde una API externa, la almacena en una base de datos CockroachDB y la presenta en una interfaz web sencilla.

---

## 🚀 Tecnologías utilizadas

| Componente | Tecnología |
|-------------|-------------|
| Backend     | Go (Gin) |
| Frontend    | Vue 3 + TypeScript + Pinia + Tailwind CSS |
| Base de datos | CockroachDB |
| Contenedores | Docker & Docker Compose |

---

## 📂 Estructura del proyecto

```
stocks-challenge/
│
├── backend/               # API y lógica del servidor
│   ├── main.go
│   ├── internal/
│   │   └── api/
│   │       └── routes.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── frontend/              # Interfaz web con Vue 3
│   ├── src/
│   │   ├── main.ts
│   │   └── App.vue
│   ├── package.json
│   ├── vite.config.ts
│   └── Dockerfile
│
├── docker-compose.yml     # Orquestación de servicios
├── .gitignore
└── README.md
```

---

## ⚙️ Configuración inicial

Clona el repositorio y entra en el proyecto:

```bash
git clone https://github.com/<tu-usuario>/stocks-challenge.git
cd stocks-challenge
```

---

## 🐳 Levantar con Docker

Ejecuta todos los servicios con un solo comando:

```bash
docker compose up --build
```

Esto levantará:

- 🧠 **Backend** en [http://localhost:8081/api/health](http://localhost:8081/api/health)
- 💻 **Frontend** en [http://localhost:5173](http://localhost:5173)
- 🗄️ **CockroachDB** (solo si se incluye más adelante)

---

## 🧪 Verificar que funciona

1. Abre en tu navegador:  
   [http://localhost:8081/api/health](http://localhost:8081/api/health)

   Debes ver algo como:
   ```json
   {
     "status": "ok",
     "service": "backend ready"
   }
   ```

2. Luego abre:  
   [http://localhost:5173](http://localhost:5173)  
   para ver la interfaz base.

---

## 🌱 Variables de entorno

Ejemplo de `.env` (no se sube al repositorio):

```bash
# Backend
PORT=8081
DATABASE_URL=<tu-connection-string>

# Frontend
VITE_API_URL=http://localhost:8081
```

> ⚠️ No incluyas tus credenciales o API keys en el repositorio público.

---

## 🧠 Cómo contribuir / probar localmente

Para ejecutar el backend sin Docker:

```bash
cd backend
go run main.go
```

Para ejecutar el frontend:

```bash
cd frontend
npm install
npm run dev
```

---

## 📄 Licencia

Proyecto creado con fines técnicos y educativos como parte de un desafío de ingeniería.  
No incluye datos sensibles ni credenciales.
