import type { MemberBirthday } from "@/types/member";

type BirthdayMemberTableProps = {
  members: MemberBirthday[];
};

export default function BirthdayMemberTable({
  members,
}: BirthdayMemberTableProps) {
  if (members.length === 0) {
    return <p>この月に誕生日のメンバーはいません。</p>;
  }

  return (
    <table>
      <thead>
        <tr>
          <th>名前</th>
          <th>誕生日</th>
        </tr>
      </thead>

      <tbody>
        {members.map((member) => (
          <tr key={`${member.name}-${member.month}-${member.date}`}>
            <td>{member.name}</td>
            <td>
              {member.month}月{member.date}日
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
