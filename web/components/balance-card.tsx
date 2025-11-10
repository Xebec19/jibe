export function BalanceCard() {
  return (
    <div className="grid gap-6 md:grid-cols-3">
      {/* Total Balance */}
      <div className="rounded-2xl bg-card p-8 shadow-sm ring-1 ring-border">
        <p className="mb-2 text-sm font-medium text-muted-foreground">
          Total Balance
        </p>
        <div className="mb-6">
          <p className="text-5xl font-bold text-foreground">$24,850</p>
          <p className="mt-1 text-sm text-accent">+12.5% this month</p>
        </div>
      </div>

      {/* Crypto Holdings */}
      <div className="rounded-2xl bg-card p-8 shadow-sm ring-1 ring-border">
        <p className="mb-2 text-sm font-medium text-muted-foreground">
          Crypto Holdings
        </p>
        <div className="mb-6">
          <p className="text-5xl font-bold text-foreground">3.24 ETH</p>
          <p className="mt-1 text-sm text-muted-foreground">≈ $12,450</p>
        </div>
      </div>

      {/* NFT Value */}
      <div className="rounded-2xl bg-card p-8 shadow-sm ring-1 ring-border">
        <p className="mb-2 text-sm font-medium text-muted-foreground">
          NFT Portfolio
        </p>
        <div className="mb-6">
          <p className="text-5xl font-bold text-foreground">12 NFTs</p>
          <p className="mt-1 text-sm text-muted-foreground">≈ $12,400</p>
        </div>
      </div>
    </div>
  );
}
