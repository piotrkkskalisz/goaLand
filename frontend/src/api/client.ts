const API_URL = import.meta.env.VITE_API_URL;

export async function get<T>(path: string): Promise<T> {
  const response = await fetch(`${API_URL}${path}`);

  if (!response.ok) {
    throw new Error(`Request failed: ${response.status}`);
  }

  const data = (await response.json()) as T;
  return data;
}