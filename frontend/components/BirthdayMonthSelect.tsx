"use client";

type BirthdayMonthSelectProps = {
  value: number | null;
  onChange: (month: number | null) => void;
};

export default function BirthdayMonthSelect({
  value,
  onChange,
}: BirthdayMonthSelectProps) {
  return (
    <label>
      誕生月：
      <select
        value={value ?? ""}
        onChange={(event) => {
          const value = event.target.value;

          onChange(value === "" ? null : Number(value));
        }}
      >
        <option value="">選択してください</option>
        {Array.from({ length: 12 }, (_, index) => {
          const month = index + 1;

          return (
            <option key={month} value={month}>
              {month}月
            </option>
          );
        })}
      </select>
    </label>
  );
}
