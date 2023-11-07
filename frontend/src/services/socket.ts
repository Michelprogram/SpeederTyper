import { EventTypes, eventListeners } from "@/services/events";
import { reactive } from "vue";

type WebSocketEvent = {
  name: string;
  data?: any;
};

export const state = reactive({
  connected: false,
});

class WebSocketService {
  private socket!: WebSocket;

  constructor(private url: string) {
    this.connect();
  }

  public connect() {
    this.socket = new WebSocket(this.url);
    this.socket.onopen = this.onOpen.bind(this);
    this.socket.onmessage = this.onMessage.bind(this);
    this.socket.onclose = this.onClose.bind(this);
    this.socket.onerror = this.onError.bind(this);
  }

  public disconnect() {
    this.socket.close();
    state.connected = false;
  }

  private onError(e: Event): void {
    console.warn("Error during connection : " + e);
  }

  private onOpen(_: Event): void {
    state.connected = true;
  }

  private onMessage(event: MessageEvent): void {
    const data = JSON.parse(event.data);
    const eventName: EventTypes = data.name;

    if (eventListeners[eventName]) {
      eventListeners[eventName](data.data);
    }
  }

  private onClose(_: CloseEvent): void {
    state.connected = false;
  }

  // Méthode pour envoyer des messages au serveur
  public sendMessage(message: WebSocketEvent): void {
    this.socket.send(JSON.stringify(message));
  }
}

const url = import.meta.env.DEV
  ? "ws://localhost:3000/ws"
  : "wss://websocket-typer-3e4343a5a6b9.herokuapp.com/ws";

let webSocketService: WebSocketService;

try {
  webSocketService = new WebSocketService(url);
} catch (e) {
  console.warn("Can't connect to the web socket");
}

export { webSocketService };
