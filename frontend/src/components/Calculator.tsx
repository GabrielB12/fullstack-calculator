import { useState } from "react";
import { calculate } from "../services/api";
import type { Operation } from "../services/api";

const OPERATIONS: { label: string; value: Operation; needsB: boolean }[] = [
  { label: "+", value: "add", needsB: true },
  { label: "−", value: "subtract", needsB: true },
  { label: "×", value: "multiply", needsB: true },
  { label: "÷", value: "divide", needsB: true },
  { label: "^", value: "power", needsB: true },
  { label: "√", value: "sqrt", needsB: false },
  { label: "%", value: "percentage", needsB: true },
];

export default function Calculator() {
  const [a, setA] = useState("");
  const [b, setB] = useState("");
  const [operation, setOperation] = useState<Operation>("add");
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const currentOp = OPERATIONS.find((o) => o.value === operation)!;

  function validate(): string | null {
    if (a.trim() === "" || isNaN(Number(a))) {
      return "Please enter a valid number for the first value.";
    }
    if (currentOp.needsB && (b.trim() === "" || isNaN(Number(b)))) {
      return "Please enter a valid number for the second value.";
    }
    return null;
  }

  async function handleCalculate() {
    setError(null);
    setResult(null);

    const validationError = validate();
    if (validationError) {
      setError(validationError);
      return;
    }

    setLoading(true);
    try {
      const numA = Number(a);
      const numB = currentOp.needsB ? Number(b) : undefined;
      const res = await calculate(operation, numA, numB);
      setResult(res);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Something went wrong.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="calculator" data-testid="calculator">
      <h1>Calculator</h1>

      <div className="field">
        <label htmlFor="value-a">Value A</label>
        <input
          id="value-a"
          type="text"
          inputMode="decimal"
          value={a}
          onChange={(e) => setA(e.target.value)}
          placeholder="e.g. 10"
        />
      </div>

      <div className="field">
        <label htmlFor="operation">Operation</label>
        <select
          id="operation"
          value={operation}
          onChange={(e) => setOperation(e.target.value as Operation)}
        >
          {OPERATIONS.map((op) => (
            <option key={op.value} value={op.value}>
              {op.label} {op.value}
            </option>
          ))}
        </select>
      </div>

      {currentOp.needsB && (
        <div className="field">
          <label htmlFor="value-b">Value B</label>
          <input
            id="value-b"
            type="text"
            inputMode="decimal"
            value={b}
            onChange={(e) => setB(e.target.value)}
            placeholder="e.g. 5"
          />
        </div>
      )}

      <button onClick={handleCalculate} disabled={loading}>
        {loading ? "Calculating..." : "Calculate"}
      </button>

      {error && <p className="error" role="alert">{error}</p>}
      {result !== null && (
        <p className="result" data-testid="result">
          Result: {result}
        </p>
      )}
    </div>
  );
}