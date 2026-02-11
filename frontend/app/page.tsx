"use client";

import { useRouter } from "next/navigation";
import { FormEvent, useCallback, useState } from "react";

function generateRoomId() {
  return Math.random().toString(36).slice(2, 8);
}

export default function HomePage() {
  const router = useRouter();
  const [roomId, setRoomId] = useState("");

  const handleCreate = useCallback(() => {
    const id = generateRoomId();
    router.push(`/room/${id}`);
  }, [router]);

  const handleJoin = useCallback(
    (e: FormEvent) => {
      e.preventDefault();
      const trimmed = roomId.trim();
      if (!trimmed) return;
      router.push(`/room/${trimmed}`);
    },
    [roomId, router],
  );

  return (
    <section className="rounded-2xl border border-divider bg-surface px-8 py-9 shadow-soft">
      <h1 className="text-xl font-semibold tracking-tight text-ink">
        Hand off a file, then leave.
      </h1>
      <p className="mt-2 max-w-md text-sm leading-relaxed text-subtle">
        Create a short-lived room, share the link, and drop a file. The relay
        just forwards bytes between browsers – nothing is stored.
      </p>

      <div className="mt-7 space-y-4">
        <button
          type="button"
          onClick={handleCreate}
          className="inline-flex w-full items-center justify-center rounded-full border border-transparent bg-ink px-4 py-2.5 text-sm font-medium text-sand transition hover:bg-ink-soft focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ink/70 focus-visible:ring-offset-2 focus-visible:ring-offset-sand"
        >
          Create a new room
        </button>

        <form
          onSubmit={handleJoin}
          className="flex items-center gap-2 rounded-full border border-divider-soft bg-soft px-3 py-2"
        >
          <input
            type="text"
            placeholder="Join by room ID"
            value={roomId}
            onChange={(e) => setRoomId(e.target.value)}
            className="flex-1 border-none bg-transparent px-2 py-1 text-sm text-ink outline-none ring-0"
          />
          <button
            type="submit"
            className="inline-flex items-center justify-center rounded-full border border-divider bg-surface px-3 py-1.5 text-xs font-medium text-ink transition hover:border-ink hover:bg-soft focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ink/70 focus-visible:ring-offset-2 focus-visible:ring-offset-soft"
          >
            Join
          </button>
        </form>
      </div>
    </section>
  );
}

