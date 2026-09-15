import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { vi, describe, it, expect, beforeEach } from "vitest";
import Calculator from "./Calculator";
import * as api from "../services/api";

describe("Calculator", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it("shows a validation error when input A is empty", async () => {
    render(<Calculator />);
    const user = userEvent.setup();
    await user.click(screen.getByText("Calculate"));
    expect(await screen.findByRole("alert")).toHaveTextContent(
      "valid number for the first value"
    );
  });

  it("calls the API and shows the result", async () => {
    vi.spyOn(api, "calculate").mockResolvedValue(8);
    render(<Calculator />);
    const user = userEvent.setup();

    await user.type(screen.getByLabelText("Value A"), "5");
    await user.type(screen.getByLabelText("Value B"), "3");
    await user.click(screen.getByText("Calculate"));

    await waitFor(() =>
      expect(screen.getByTestId("result")).toHaveTextContent("Result: 8")
    );
  });

  it("shows an API error message (e.g. division by zero)", async () => {
    vi.spyOn(api, "calculate").mockRejectedValue(new Error("division by zero"));
    render(<Calculator />);
    const user = userEvent.setup();

    await user.type(screen.getByLabelText("Value A"), "5");
    await user.type(screen.getByLabelText("Value B"), "0");
    await user.click(screen.getByText("Calculate"));

    expect(await screen.findByRole("alert")).toHaveTextContent(
      "division by zero"
    );
  });

  it("hides Value B for sqrt operation", async () => {
    render(<Calculator />);
    const user = userEvent.setup();
    await user.selectOptions(screen.getByLabelText("Operation"), "sqrt");
    expect(screen.queryByLabelText("Value B")).not.toBeInTheDocument();
  });
});