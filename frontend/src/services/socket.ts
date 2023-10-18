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
  private socket: WebSocket;

  constructor(private url: string) {
    this.socket = new WebSocket(url);
    this.socket.onopen = this.onOpen.bind(this);
    this.socket.onmessage = this.onMessage.bind(this);
    this.socket.onclose = this.onClose.bind(this);
    state.connected = true;
  }

  public connect() {
    this.socket = new WebSocket(this.url);
    this.socket.onopen = this.onOpen.bind(this);
    this.socket.onmessage = this.onMessage.bind(this);
    this.socket.onclose = this.onClose.bind(this);
    state.connected = true;
  }

  public disconnect() {
    this.socket.close();
    state.connected = false;
  }

  private onOpen(_: Event): void {
    console.log("Connected to " + this.url);
  }

  private onMessage(event: MessageEvent): void {
    const data = JSON.parse(event.data);
    const eventName: EventTypes = data.name;

    if (eventListeners[eventName]) {
      eventListeners[eventName](data.data);
    }
  }

  private onClose(_: CloseEvent): void {
    console.log("Disconnected to " + this.url);
  }

  // Méthode pour envoyer des messages au serveur
  public sendMessage(message: WebSocketEvent): void {
    this.socket.send(JSON.stringify(message));
  }
}

const URL =
  window.location.hostname == "localhost"
    ? "ws://localhost:3000/ws"
    : "ws://10.101.0.122:3000/ws";

export const webSocketService = new WebSocketService(URL);
