import type { ReactNode } from "react";
import "./globals.css";

export const metadata = {
  title: "holo – ephemeral file sharing",
  description: "Minimal peer-to-peer–style file sharing over WebSockets.",
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body className="min-h-screen bg-sand text-ink antialiased">
        <div className="mx-auto flex min-h-screen max-w-3xl flex-col px-5 py-8">
          <header className="flex items-center justify-between pb-6">
            <div>
              <div className="text-xs font-medium tracking-[0.18em] text-muted uppercase">
                holo
              </div>
              <div className="mt-1 text-xs text-subtle">
                ephemeral, room-based file handoff
              </div>
            </div>
          </header>
          <main className="flex flex-1 items-center justify-center">
            <div className="w-full max-w-xl">{children}</div>
          </main>
          <footer className="mt-6 border-t border-divider pt-3 text-[11px] text-subtle">
            no accounts · no history · just a room and a file
          </footer>
        </div>
      </body>
    </html>
  );
}

