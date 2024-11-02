// app/layout.js or app/layout.tsx (depending on your setup)
import "./globals.css";

export const metadata = {
  title: "Lemanspedia - WikiRace Program",
  description: "WikiRace Program with BFS and IDS Algorithms",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className="bg-[#f7f8fa]">
        {children}
      </body>
    </html>
  );
}
