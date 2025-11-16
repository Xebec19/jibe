'use client'

import { Lock, Wallet, Users, BarChart3, Zap, Shield } from 'lucide-react'

export function Features() {
  const features = [
    {
      icon: Lock,
      title: 'Token Gating',
      description: 'Gate content behind NFT or token ownership. Your rules, your community.',
    },
    {
      icon: Wallet,
      title: 'Direct Payments',
      description: 'Earn crypto instantly. Money goes straight to your wallet.',
    },
    {
      icon: Users,
      title: 'Community First',
      description: 'Tools to engage members, build relationships, and grow organically.',
    },
    {
      icon: BarChart3,
      title: 'Analytics',
      description: 'Understand who your audience is and what content resonates.',
    },
    {
      icon: Zap,
      title: 'Multi-chain',
      description: 'Works with Ethereum, Polygon, Base, and more blockchains.',
    },
    {
      icon: Shield,
      title: 'Secure',
      description: 'Audited smart contracts. Your content and earnings are safe.',
    },
  ]

  return (
    <section id="features" className="py-24 px-6">
      <div className="max-w-6xl mx-auto">
        <div className="text-center space-y-3 mb-16">
          <h2 className="text-4xl font-bold">Everything you need</h2>
          <p className="text-lg text-muted-foreground">
            Purpose-built for creators who want to own their community.
          </p>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          {features.map((feature, i) => {
            const Icon = feature.icon
            return (
              <div
                key={i}
                className="space-y-3"
              >
                <div className="w-10 h-10 rounded-lg bg-muted flex items-center justify-center">
                  <Icon size={20} className="text-primary" />
                </div>
                <h3 className="font-bold text-lg">{feature.title}</h3>
                <p className="text-muted-foreground text-sm leading-relaxed">
                  {feature.description}
                </p>
              </div>
            )
          })}
        </div>
      </div>
    </section>
  )
}
