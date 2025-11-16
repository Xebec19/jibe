'use client'

import { Button } from '@/components/ui/button'
import { ArrowRight } from 'lucide-react'

export function Hero() {
  return (
    <section className="min-h-screen flex items-center justify-center px-6 py-20">
      <div className="max-w-3xl mx-auto text-center space-y-8">
        
        {/* Main Heading */}
        <div className="space-y-6">
          <h1 className="text-5xl md:text-6xl font-bold leading-tight text-balance">
            Build your{' '}
            <span className="text-primary">token-gated</span>{' '}
            community
          </h1>
          <p className="text-lg md:text-xl text-muted-foreground text-balance max-w-2xl mx-auto leading-relaxed">
            Share exclusive content with people who own your NFT or token. Direct access. Direct payment. No middleman.
          </p>
        </div>

        {/* CTA Buttons */}
        <div className="flex flex-col sm:flex-row gap-3 justify-center pt-4">
          <Button className="bg-primary text-primary-foreground hover:bg-primary/90 px-6 h-11 text-base font-medium">
            Get Started
            <ArrowRight className="ml-2" size={18} />
          </Button>
          <Button
            variant="outline"
            className="border-foreground/20 px-6 h-11 text-base font-medium hover:bg-muted"
          >
            See Examples
          </Button>
        </div>

        {/* Social proof - simple, no flourish */}
        <div className="pt-8 border-t border-border">
          <p className="text-xs text-muted-foreground mb-4">Trusted by creators</p>
          <div className="flex justify-center gap-8 text-sm font-medium text-foreground">
            <div>1,200+ creators</div>
            <div>$8M+ distributed</div>
            <div>50K+ members</div>
          </div>
        </div>
      </div>
    </section>
  )
}
