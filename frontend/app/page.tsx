"use client";

const apiBaseUrl =
  process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080/api";
const backendUrl = apiBaseUrl.replace(/\/api\/?$/, "");

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-[#161a1d] px-6 text-white">
      <section className="w-full max-w-xl border border-[#3d454b] bg-[#22272b] p-8 shadow-2xl sm:p-12">
        <p className="mb-8 text-sm font-semibold uppercase tracking-[0.25em] text-[#9ca8b1]">
          235bot management
        </p>
        <h1 className="max-w-md text-4xl font-semibold leading-tight sm:text-5xl">
          235botを、もっと身近に。
        </h1>
        <p className="mt-6 max-w-md leading-7 text-[#c6cdd2]">
          Discordサーバーのメンバーだけが利用できる管理サイトです。
        </p>
        <a
          className="mt-10 inline-flex h-12 items-center justify-center bg-[#5865f2] px-6 font-semibold transition-colors hover:bg-[#4752c4]"
          href={`${backendUrl}/discord/auth`}
        >
          Discordでログイン
        </a>
      </section>
    </main>
  );
}
