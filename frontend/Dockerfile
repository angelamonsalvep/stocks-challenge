# Etapa base
FROM node:20-alpine

WORKDIR /app

# Copiar dependencias
COPY package*.json ./
RUN npm install

# Copiar el resto del código
COPY . .

# Exponer el puerto de Vite
EXPOSE 5173

# Comando por defecto
CMD ["npm", "run", "dev", "--", "--host"]
