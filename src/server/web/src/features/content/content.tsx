import type { JSX, ReactNode } from "react";

interface ContentProps {
  children: ReactNode;
}

export const Content = ({ children }: ContentProps): JSX.Element | null => {
  return (
    <main className="body-container">
      {children}
    </main>
  );
};