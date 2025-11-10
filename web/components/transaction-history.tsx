const transactions = [
  {
    id: 1,
    type: "received",
    label: "Received ETH",
    amount: "+0.5 ETH",
    timestamp: "2 hours ago",
    status: "completed",
  },
  {
    id: 2,
    type: "sent",
    label: "Sent to wallet",
    amount: "-0.2 ETH",
    timestamp: "5 hours ago",
    status: "completed",
  },
  {
    id: 3,
    type: "nft",
    label: "NFT Purchase",
    amount: "-2.1 ETH",
    timestamp: "1 day ago",
    status: "completed",
  },
  {
    id: 4,
    type: "received",
    label: "Staking Reward",
    amount: "+0.05 ETH",
    timestamp: "2 days ago",
    status: "completed",
  },
  {
    id: 5,
    type: "sent",
    label: "Bridge Transfer",
    amount: "-1.0 ETH",
    timestamp: "3 days ago",
    status: "pending",
  },
]

export function TransactionHistory() {
  return (
    <div>
      <div className="mb-8">
        <h2 className="text-2xl font-bold text-foreground">Recent Activity</h2>
      </div>

      <div className="space-y-3">
        {transactions.map((tx) => (
          <div key={tx.id} className="rounded-lg bg-card p-4 ring-1 ring-border transition-colors hover:bg-secondary">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-3">
                <div
                  className={`h-10 w-10 rounded-full flex items-center justify-center text-sm font-semibold ${
                    tx.type === "received"
                      ? "bg-accent/10 text-accent"
                      : tx.type === "nft"
                        ? "bg-primary/10 text-primary"
                        : "bg-muted text-muted-foreground"
                  }`}
                >
                  {tx.type === "received" ? "↓" : tx.type === "nft" ? "◆" : "↑"}
                </div>
                <div>
                  <p className="font-medium text-foreground text-sm">{tx.label}</p>
                  <p className="text-xs text-muted-foreground">{tx.timestamp}</p>
                </div>
              </div>
              <div className="text-right">
                <p className={`font-semibold text-sm ${tx.type === "received" ? "text-accent" : "text-foreground"}`}>
                  {tx.amount}
                </p>
                <p className={`text-xs ${tx.status === "completed" ? "text-accent" : "text-yellow-600"}`}>
                  {tx.status === "completed" ? "Done" : "Pending"}
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>

      <button className="mt-8 w-full rounded-lg border border-border bg-card py-2 text-sm font-medium text-foreground transition-colors hover:bg-secondary">
        View All Transactions
      </button>
    </div>
  )
}
