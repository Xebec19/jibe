import { ConnectButton } from "@rainbow-me/rainbowkit";
import Link from "next/link";

export function Navbar() {
  return (
    <nav className="border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo/Brand */}
          <div className="flex-shrink-0">
            <Link
              href="/"
              className="text-xl font-bold text-foreground hover:text-accent transition-colors"
            >
              Web3 Profile
            </Link>
          </div>

          {/* Navigation Links */}
          <div className="hidden md:flex gap-8">
            <Link
              href="/"
              className="text-sm text-foreground/70 hover:text-foreground transition-colors"
            >
              Profile
            </Link>
            <Link
              href="/portfolio"
              className="text-sm text-foreground/70 hover:text-foreground transition-colors"
            >
              Portfolio
            </Link>
            <Link
              href="/activity"
              className="text-sm text-foreground/70 hover:text-foreground transition-colors"
            >
              Activity
            </Link>
            <Link
              href="/settings"
              className="text-sm text-foreground/70 hover:text-foreground transition-colors"
            >
              Settings
            </Link>
          </div>

          {/* Right Side - Connect Wallet Button */}
          {/* <div className="flex items-center gap-4">
            <button className="px-4 py-2 text-sm font-medium text-foreground-foreground bg-primary text-primary-foreground rounded-lg hover:opacity-90 transition-opacity">
              Connect Wallet
            </button>
          </div> */}
          <ConnectButton />
        </div>
      </div>
    </nav>
  );
}
