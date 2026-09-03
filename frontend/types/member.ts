export type MemberBirthday = {
  name: string;
  month: number;
  date: number;
};

export type GetMembersBirthdayResponse = {
  result: string;
  members: MemberBirthday[];
};
