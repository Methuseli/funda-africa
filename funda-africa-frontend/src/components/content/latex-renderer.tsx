"use client"

import React from 'react'
import 'katex/dist/katex.min.css'
import { InlineMath, BlockMath } from 'react-katex'

interface LaTeXRendererProps {
  content: string
  block?: boolean
  className?: string
}

export default function LaTeXRenderer({ content, block = false, className = '' }: LaTeXRendererProps) {
  try {
    if (block) {
      return (
        <div className={`my-4 ${className}`}>
          <BlockMath math={content} />
        </div>
      )
    } else {
      return <InlineMath math={content} />
    }
  } catch (error) {
    console.error('LaTeX rendering error:', error)
    return (
      <span className={`font-mono text-red-600 bg-red-50 px-1 rounded ${className}`}>
        LaTeX Error: {content}
      </span>
    )
  }
}