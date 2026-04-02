#!/bin/bash

# --- CONFIGURATION ---
BINARY_NAME="download-tool"
SERVER_USER="mihu"
SERVER_IP="192.168.0.201"
DEST_DIR="/mnt/containers/downloads"
COMPOSE_DIR="/mnt/containers"

echo "🚀 Starte Deployment für $BINARY_NAME..."

# 1. Kompilieren
echo "📦 Kompiliere Go-Projekt..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $BINARY_NAME .

if [ $? -ne 0 ]; then
    echo "❌ Fehler beim Kompilieren!"
    exit 1
fi

# 2. Container stoppen (Using the SERVICE name 'download-tool')
echo "🛑 Stoppe Downloads-Service..."
ssh -t $SERVER_USER@$SERVER_IP "cd $COMPOSE_DIR && podman-compose stop download-tool"

# 3. Datei übertragen
echo "🚚 Übertrage $BINARY_NAME nach $SERVER_IP..."
scp $BINARY_NAME $SERVER_USER@$SERVER_IP:$DEST_DIR/

if [ $? -ne 0 ]; then
    echo "❌ Fehler bei der Übertragung!"
    exit 1
fi

# 4. Rechte setzen und Container wieder starten
echo "🔄 Setze Rechte und starte Service neu..."
ssh -t $SERVER_USER@$SERVER_IP "
    chmod +x $DEST_DIR/$BINARY_NAME && \
    cd $COMPOSE_DIR && \
    podman-compose up -d download-tool && \
    echo '✅ Download-Service erfolgreich aktualisiert.'
"

# 5. Live Logs anzeigen
echo "📋 Zeige Live-Logs von 'downloads' (Beenden mit Ctrl+C)..."
ssh -t $SERVER_USER@$SERVER_IP "podman logs -f downloads"

echo "🎉 Deployment beendet!"