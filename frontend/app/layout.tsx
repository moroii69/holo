import type { ReactNode } from "react";
import "./globals.css";

export const metadata = {
  title: "holo – ephemeral file sharing",
  description: "Minimal peer-to-peer–style file sharing over WebSockets.",
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body className="min-h-screen bg-sand text-ink">
        <div className="mx-auto flex min-h-screen max-w-2xl flex-col px-5 py-8 sm:px-6 sm:py-12">
          <header className="mb-8 sm:mb-12">
            <div className="text-xs font-medium tracking-[0.2em] text-muted uppercase">
              holo
            </div>
          </header>
          <main className="flex flex-1 items-center justify-center">
            <div className="w-full">{children}</div>
          </main>
          <footer className="mt-auto pt-12 text-[11px] text-muted">
            no storage · no accounts · just a room
          </footer>
        </div>
      </body>
    </html>
  );
}

