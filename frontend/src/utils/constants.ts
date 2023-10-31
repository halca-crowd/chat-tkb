export const randomStr = Math.random().toString(32).substring(2)
export const DEFAULT_OPTIONS = {}
export const DEFAULT_RECONNECT_LIMIT = 20
export const DEFAULT_RECONNECT_INTERVAL_MS = 5000
export const UNPARSABLE_JSON_OBJECT = {}
export const ACTION_RECV_MESSAGE = 'gpt_message'
export const ACTION_SEND_MESSAGE = 'chat_message'
export const ACTION_RECV_STATUS = 'ACTION_SEND_STATUS'
export const ACTION_SEND_MASAKARI = 'gpt_message'
export const ACTION_RECV_MASAKARI = 'ACTION_GPT_MESSAGE'
export const ACTION_FORCE_RESET = 'reset_message'
export enum ReadyState {
  UNINSTANTIATED = -1,
  CONNECTING = 0,
  OPEN = 1,
  CLOSING = 2,
  CLOSED = 3,
}
export const BASE_API_URL =
  import.meta.env.VITE_API_URL || 'http://localhost:8080'
const eventSourceSupported = () => {
  try {
    return 'EventSource' in globalThis
  } catch (e) {
    return false
  }
}

export const isEventSourceSupported = eventSourceSupported()
