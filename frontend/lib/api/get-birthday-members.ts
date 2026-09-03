import { GetMembersBirthdayResponse, MemberBirthday } from "@/types/member";

export async function getBirthdayMembers(
  month: number,
): Promise<MemberBirthday[]> {
  const API_BASE_URL =
    process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080/api";
  const response = await fetch(
    `${API_BASE_URL}/members?birthday_month=${month}`,
  );

  if (!response.ok) {
    throw new Error("誕生日メンバーの取得に失敗しました");
  }

  const data: GetMembersBirthdayResponse = await response.json();

  return data.members;
}
