const nfts = [
  {
    id: 1,
    title: "Digital Canvas #427",
    image: "/digital-art-nft-purple.jpg",
    floor: "2.5 ETH",
    rarity: "Rare",
  },
  {
    id: 2,
    title: "Pixel Dreams",
    image: "/pixel-art-nft-colorful.jpg",
    floor: "1.8 ETH",
    rarity: "Common",
  },
  {
    id: 3,
    title: "Abstract Genesis",
    image: "/abstract-nft-blue.jpg",
    floor: "3.2 ETH",
    rarity: "Legendary",
  },
  {
    id: 4,
    title: "Cyber Punk #89",
    image: "/cyberpunk-nft-digital.jpg",
    floor: "1.4 ETH",
    rarity: "Uncommon",
  },
];

export function NFTGrid() {
  return (
    <div>
      <div className="mb-12">
        <h2 className="text-2xl font-bold text-foreground">NFTs Owned</h2>
        <p className="mt-1 text-sm text-muted-foreground">
          12 items in collection
        </p>
      </div>

      <div className="grid gap-6 sm:grid-cols-2">
        {nfts.map((nft) => (
          <div
            key={nft.id}
            className="group cursor-pointer rounded-xl bg-card overflow-hidden shadow-sm ring-1 ring-border transition-all hover:shadow-md hover:ring-accent/50"
          >
            <div className="aspect-square overflow-hidden bg-secondary">
              <img
                src={nft.image || "/placeholder.svg"}
                alt={nft.title}
                className="h-full w-full object-cover transition-transform group-hover:scale-105"
              />
            </div>
            <div className="p-6">
              <h3 className="font-bold text-foreground">{nft.title}</h3>
              <div className="mt-4 flex items-center justify-between">
                <div>
                  <p className="text-xs text-muted-foreground">Floor Price</p>
                  <p className="mt-1 font-semibold text-foreground">
                    {nft.floor}
                  </p>
                </div>
                <span className="rounded-full bg-accent/10 px-3 py-1 text-xs font-medium text-accent">
                  {nft.rarity}
                </span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
