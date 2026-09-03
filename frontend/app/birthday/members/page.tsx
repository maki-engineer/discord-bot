"use client";

import { useEffect, useState } from "react";

import BirthdayMemberTable from "@/components/BirthdayMemberTable";
import BirthdayMonthSelect from "@/components/BirthdayMonthSelect";
import type { MemberBirthday } from "@/types/member";
import { getBirthdayMembers } from "@/lib/api/get-birthday-members";

export default function BirthdayMembersPage() {
  const [month, setMonth] = useState<number | null>(null);
  const [members, setMembers] = useState<MemberBirthday[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (month === null) {
      return;
    }

    const fetchMembers = async () => {
      setIsLoading(true);
      setError(null);

      try {
        const data = await getBirthdayMembers(month);
        setMembers(data);
      } catch {
        setError("誕生日メンバーの取得に失敗しました");
        setMembers([]);
      } finally {
        setIsLoading(false);
      }
    };

    void fetchMembers();
  }, [month]);

  return (
    <main>
      <h1>誕生日メンバー</h1>

      <BirthdayMonthSelect value={month} onChange={setMonth} />

      {isLoading && <p>読み込み中...</p>}

      {error && <p>{error}</p>}

      {month !== null && !isLoading && !error && (
        <BirthdayMemberTable members={members} />
      )}
    </main>
  );
}
