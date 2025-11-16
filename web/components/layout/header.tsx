"use client";

import Link from "next/link";
import { useState } from "react";
import { Menu, X } from "lucide-react";
import { Button } from "@/components/ui/button";

export function Header() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <header className="sticky top-0 z-50 bg-background/95 backdrop-blur-sm border-b border-border">
      <nav className="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 bg-primary rounded-lg flex items-center justify-center">
            <span className="text-primary-foreground font-bold">◈</span>
          </div>
          <span className="font-bold text-lg">TokenFlow</span>
        </div>

        <div className="hidden md:flex items-center gap-8">
          <Link
            href="#market"
            className="text-sm hover:text-primary transition-colors"
          >
            Market
          </Link>
          <Link
            href="#profile"
            className="text-sm hover:text-primary transition-colors"
          >
            Profile
          </Link>
          <Link
            href="#"
            className="text-sm hover:text-primary transition-colors"
          >
            Settings
          </Link>
        </div>

        <div className="hidden md:flex items-center gap-3">
          <Button className="border-foreground/20">Connect Wallet</Button>
        </div>

        <button
          className="md:hidden"
          onClick={() => setIsOpen(!isOpen)}
          aria-label="Toggle menu"
        >
          {isOpen ? <X size={24} /> : <Menu size={24} />}
        </button>

        {isOpen && (
          <div className="absolute top-16 left-0 right-0 bg-background border-b border-border md:hidden">
            <div className="flex flex-col gap-4 p-6">
              <Link
                href="#market"
                className="text-sm hover:text-primary transition-colors"
              >
                Market
              </Link>
              <Link
                href="#profile"
                className="text-sm hover:text-primary transition-colors"
              >
                Profile
              </Link>
              <Link
                href="#"
                className="text-sm hover:text-primary transition-colors"
              >
                Settings
              </Link>
              <div className="flex flex-col gap-2 pt-4 border-t border-border">
                <Button className="w-full border-foreground/20">
                  Connect Wallet
                </Button>
              </div>
            </div>
          </div>
        )}
      </nav>
    </header>
  );
}
