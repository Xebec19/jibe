"use client"

import { useWallet } from "@/contexts/wallet-context"

export function WalletNotConnected() {
  const { connect } = useWallet()

  return (
    <main className="min-h-screen bg-background flex items-center justify-center px-4">
      <div className="w-full max-w-md text-center">
        {/* Icon */}
        <div className="mb-8 flex justify-center">
          <div className="h-24 w-24 rounded-full bg-accent/10 flex items-center justify-center">
            <svg className="h-12 w-12 text-accent" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={1.5}
                d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.658 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"
              />
            </svg>
          </div>
        </div>

        {/* Heading */}
        <h1 className="text-4xl font-bold text-foreground mb-3">Connect Your Wallet</h1>

        {/* Description */}
        <p className="text-lg text-muted-foreground mb-8 leading-relaxed">
          Connect your Web3 wallet to view your balance, NFTs, and transaction history.
        </p>

        {/* Connect Button */}
        <button
          onClick={connect}
          className="w-full rounded-lg bg-primary px-8 py-4 text-lg font-semibold text-primary-foreground transition-all hover:opacity-90 active:scale-95"
        >
          Connect Wallet
        </button>

        {/* Info Text */}
        <p className="mt-8 text-sm text-muted-foreground">
          Supported wallets: MetaMask, WalletConnect, Ledger, and more
        </p>
      </div>
    </main>
  )
}
