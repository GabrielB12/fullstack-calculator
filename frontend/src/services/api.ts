const BASE_URL = "http://localhost:8080/api";

export type Operation =
  | "add"
  | "subtract"
  | "multiply"
  | "divide"
  | "power"
  | "sqrt"
  | "percentage";

interface CalcResponse {
  result: number;
}

interface ErrorResponse {
  error: string;
}

export async function calculate(
  operation: Operation,
  a: number,
  b?: number
): Promise<number> {
  const body = b === undefined ? { a } : { a, b };

  const response = await fetch(`${BASE_URL}/${operation}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });

  const data = await response.json();

  if (!response.ok) {
    const err = data as ErrorResponse;
    throw new Error(err.error || "Unknown error");
  }

  return (data as CalcResponse).result;
}