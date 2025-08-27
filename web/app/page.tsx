"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  Wallet,
  Shield,
  Zap,
  Users,
  Star,
  ChevronRight,
  ExternalLink,
  Copy,
  Check,
} from "lucide-react";

export default function Web3ContentPlatform() {
  const [isWalletConnected, setIsWalletConnected] = useState(false);
  const [walletAddress, setWalletAddress] = useState("");
  const [copiedAddress, setCopiedAddress] = useState(false);

  const connectWallet = () => {
    // Simulate wallet connection
    setIsWalletConnected(true);
    setWalletAddress("0x742d35Cc6634C0532925a3b8D4C0532925a3b8D4");
  };

  const copyAddress = () => {
    navigator.clipboard.writeText(walletAddress);
    setCopiedAddress(true);
    setTimeout(() => setCopiedAddress(false), 2000);
  };

  const truncateAddress = (address: string) => {
    return `${address.slice(0, 6)}...${address.slice(-4)}`;
  };

  return (
    <div className="min-h-screen bg-[#0a0a0a] text-white">
      {/* Navigation */}
      <nav className="border-b border-[#2a2a2a] bg-[#0a0a0a]/80 backdrop-blur-md sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-gradient-to-br from-[#00D4FF] to-[#8B5CF6] rounded-lg flex items-center justify-center">
                <div className="w-4 h-4 bg-white rounded-sm"></div>
              </div>
              <span className="text-xl font-bold bg-gradient-to-r from-[#00D4FF] to-[#8B5CF6] bg-clip-text text-transparent">
                Jibe
              </span>
            </div>

            <div className="flex items-center space-x-6">
              <a
                href="#"
                className="text-gray-300 hover:text-[#00D4FF] transition-colors"
              >
                Features
              </a>
              <a
                href="#"
                className="text-gray-300 hover:text-[#00D4FF] transition-colors"
              >
                Creators
              </a>
              <a
                href="#"
                className="text-gray-300 hover:text-[#00D4FF] transition-colors"
              >
                Docs
              </a>

              {isWalletConnected ? (
                <div className="flex items-center space-x-2 bg-[#1a1a1a] border border-[#2a2a2a] rounded-lg px-3 py-2">
                  <div className="w-2 h-2 bg-[#10B981] rounded-full animate-pulse"></div>
                  <span className="text-sm font-mono">
                    {truncateAddress(walletAddress)}
                  </span>
                  <button
                    onClick={copyAddress}
                    className="text-gray-400 hover:text-white"
                  >
                    {copiedAddress ? (
                      <Check className="w-4 h-4" />
                    ) : (
                      <Copy className="w-4 h-4" />
                    )}
                  </button>
                </div>
              ) : (
                <Button
                  onClick={connectWallet}
                  className="bg-gradient-to-r from-[#00D4FF] to-[#8B5CF6] hover:from-[#00B8E6] hover:to-[#7C3AED] text-white border-0 shadow-lg shadow-[#00D4FF]/20"
                >
                  <Wallet className="w-4 h-4 mr-2" />
                  Connect Wallet
                </Button>
              )}
            </div>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-[#00D4FF]/10 via-transparent to-[#8B5CF6]/10"></div>
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24">
          <div className="text-center">
            <h1 className="text-5xl md:text-7xl font-bold mb-6">
              <span className="bg-gradient-to-r from-[#00D4FF] via-[#8B5CF6] to-[#10B981] bg-clip-text text-transparent">
                Decentralized
              </span>
              <br />
              Content Creation
            </h1>
            <p className="text-xl text-gray-300 mb-8 max-w-3xl mx-auto">
              Empower creators with true ownership. Build your community,
              monetize your content, and earn directly from your fans without
              intermediaries.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Button
                size="lg"
                className="bg-gradient-to-r from-[#00D4FF] to-[#8B5CF6] hover:from-[#00B8E6] hover:to-[#7C3AED] text-white border-0 shadow-lg shadow-[#00D4FF]/20"
              >
                Start Creating
                <ChevronRight className="w-5 h-5 ml-2" />
              </Button>
              <Button
                size="lg"
                variant="outline"
                className="border-[#2a2a2a] text-white hover:bg-[#1a1a1a] bg-transparent"
              >
                Explore Creators
                <ExternalLink className="w-5 h-5 ml-2" />
              </Button>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-24 bg-[#1a1a1a]/50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold mb-4">Why Choose Jibe?</h2>
            <p className="text-gray-300 text-lg">
              Built for the future of content creation
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            <Card className="bg-[#1a1a1a] border-[#2a2a2a] hover:border-[#00D4FF]/50 transition-all duration-300 group">
              <CardHeader>
                <div className="w-12 h-12 bg-gradient-to-br from-[#00D4FF] to-[#8B5CF6] rounded-lg flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                  <Shield className="w-6 h-6 text-white" />
                </div>
                <CardTitle className="text-white">True Ownership</CardTitle>
                <CardDescription className="text-gray-400">
                  Your content, your rules. Smart contracts ensure you maintain
                  full control and ownership.
                </CardDescription>
              </CardHeader>
            </Card>

            <Card className="bg-[#1a1a1a] border-[#2a2a2a] hover:border-[#8B5CF6]/50 transition-all duration-300 group">
              <CardHeader>
                <div className="w-12 h-12 bg-gradient-to-br from-[#8B5CF6] to-[#10B981] rounded-lg flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                  <Zap className="w-6 h-6 text-white" />
                </div>
                <CardTitle className="text-white">Instant Payments</CardTitle>
                <CardDescription className="text-gray-400">
                  Get paid instantly in crypto. No waiting periods, no middlemen
                  taking cuts.
                </CardDescription>
              </CardHeader>
            </Card>

            <Card className="bg-[#1a1a1a] border-[#2a2a2a] hover:border-[#10B981]/50 transition-all duration-300 group">
              <CardHeader>
                <div className="w-12 h-12 bg-gradient-to-br from-[#10B981] to-[#00D4FF] rounded-lg flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                  <Users className="w-6 h-6 text-white" />
                </div>
                <CardTitle className="text-white">Community Driven</CardTitle>
                <CardDescription className="text-gray-400">
                  Build deeper connections with token-gated content and
                  exclusive NFT rewards.
                </CardDescription>
              </CardHeader>
            </Card>
          </div>
        </div>
      </section>

      {/* Featured Creators */}
      <section className="py-24">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold mb-4">Featured Creators</h2>
            <p className="text-gray-300 text-lg">
              Join thousands of creators earning on-chain
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            {[
              {
                name: "Alex Chen",
                category: "Digital Art",
                earnings: "45.2 ETH",
                subscribers: "12.4K",
                avatar: "/digital-artist-avatar.png",
              },
              {
                name: "Sarah Kim",
                category: "Music",
                earnings: "32.8 ETH",
                subscribers: "8.9K",
                avatar: "/musician-avatar.png",
              },
              {
                name: "Mike Torres",
                category: "Gaming",
                earnings: "28.1 ETH",
                subscribers: "15.2K",
                avatar: "/gamer-avatar.png",
              },
            ].map((creator, index) => (
              <Card
                key={index}
                className="bg-[#1a1a1a] border-[#2a2a2a] hover:border-[#00D4FF]/30 transition-all duration-300 group cursor-pointer"
              >
                <CardHeader>
                  <div className="flex items-center space-x-4">
                    <Avatar className="w-16 h-16 border-2 border-[#00D4FF]/20">
                      <AvatarImage src={creator.avatar || "/placeholder.svg"} />
                      <AvatarFallback className="bg-gradient-to-br from-[#00D4FF] to-[#8B5CF6] text-white">
                        {creator.name
                          .split(" ")
                          .map((n) => n[0])
                          .join("")}
                      </AvatarFallback>
                    </Avatar>
                    <div>
                      <h3 className="text-white font-semibold">
                        {creator.name}
                      </h3>
                      <Badge
                        variant="secondary"
                        className="bg-[#2a2a2a] text-[#00D4FF] border-0"
                      >
                        {creator.category}
                      </Badge>
                    </div>
                  </div>
                </CardHeader>
                <CardContent>
                  <div className="flex justify-between items-center">
                    <div>
                      <p className="text-sm text-gray-400">Total Earned</p>
                      <p className="text-lg font-bold text-[#10B981]">
                        {creator.earnings}
                      </p>
                    </div>
                    <div className="text-right">
                      <p className="text-sm text-gray-400">Subscribers</p>
                      <p className="text-lg font-bold text-white">
                        {creator.subscribers}
                      </p>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* Stats Section */}
      <section className="py-24 bg-[#1a1a1a]/50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-4 gap-8 text-center">
            <div>
              <div className="text-4xl font-bold text-[#00D4FF] mb-2">
                $2.4M+
              </div>
              <div className="text-gray-300">Creator Earnings</div>
            </div>
            <div>
              <div className="text-4xl font-bold text-[#8B5CF6] mb-2">15K+</div>
              <div className="text-gray-300">Active Creators</div>
            </div>
            <div>
              <div className="text-4xl font-bold text-[#10B981] mb-2">
                250K+
              </div>
              <div className="text-gray-300">Supporters</div>
            </div>
            <div>
              <div className="text-4xl font-bold text-[#00D4FF] mb-2">
                99.9%
              </div>
              <div className="text-gray-300">Uptime</div>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-24">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-4xl font-bold mb-6">
            Ready to Own Your Content?
          </h2>
          <p className="text-xl text-gray-300 mb-8">
            Join the decentralized creator economy and start earning from day
            one.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button
              size="lg"
              className="bg-gradient-to-r from-[#00D4FF] to-[#8B5CF6] hover:from-[#00B8E6] hover:to-[#7C3AED] text-white border-0 shadow-lg shadow-[#00D4FF]/20"
            >
              Launch Your Channel
              <Star className="w-5 h-5 ml-2" />
            </Button>
            <Button
              size="lg"
              variant="outline"
              className="border-[#2a2a2a] text-white hover:bg-[#1a1a1a] bg-transparent"
            >
              Read Documentation
            </Button>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-[#2a2a2a] bg-[#1a1a1a]/50 py-12">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-4 gap-8">
            <div>
              <div className="flex items-center space-x-2 mb-4">
                <div className="w-8 h-8 bg-gradient-to-br from-[#00D4FF] to-[#8B5CF6] rounded-lg flex items-center justify-center">
                  <div className="w-4 h-4 bg-white rounded-sm"></div>
                </div>
                <span className="text-xl font-bold bg-gradient-to-r from-[#00D4FF] to-[#8B5CF6] bg-clip-text text-transparent">
                  Jibe
                </span>
              </div>
              <p className="text-gray-400">
                The future of decentralized content creation.
              </p>
            </div>

            <div>
              <h3 className="text-white font-semibold mb-4">Platform</h3>
              <ul className="space-y-2 text-gray-400">
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    Features
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    Pricing
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    Security
                  </a>
                </li>
              </ul>
            </div>

            <div>
              <h3 className="text-white font-semibold mb-4">Developers</h3>
              <ul className="space-y-2 text-gray-400">
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    API Docs
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    Smart Contracts
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    GitHub
                  </a>
                </li>
              </ul>
            </div>

            <div>
              <h3 className="text-white font-semibold mb-4">Community</h3>
              <ul className="space-y-2 text-gray-400">
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    Discord
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    Twitter
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00D4FF] transition-colors"
                  >
                    Blog
                  </a>
                </li>
              </ul>
            </div>
          </div>

          <div className="border-t border-[#2a2a2a] mt-8 pt-8 text-center text-gray-400">
            <p>&copy; 2024 Jibe. Built on the blockchain for creators.</p>
          </div>
        </div>
      </footer>
    </div>
  );
}
