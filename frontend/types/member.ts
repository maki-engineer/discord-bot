export type MemberBirthday = {
  name: string;
  month: number;
  date: number;
};

// テスト

export type GetMembersBirthdayResponse = {
  result: string;
  members: MemberBirthday[];
};
