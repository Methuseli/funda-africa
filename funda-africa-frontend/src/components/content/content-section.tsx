"use client"

import React from 'react'
import { ContentSection } from '@/lib/data'
import LaTeXRenderer from './latex-renderer'
import CodeBlock from './code-block'
import Image from 'next/image'

interface ContentSectionProps {
  section: ContentSection
}

export default function ContentSectionComponent({ section }: ContentSectionProps) {
  const renderContent = () => {
    switch (section.type) {
      case 'text':
        return (
          <div className="prose prose-gray max-w-none">
            <p className="text-gray-700 leading-relaxed">{section.content}</p>
          </div>
        )
      
      case 'latex':
        return (
          <LaTeXRenderer 
            content={section.content} 
            block={true}
            className="bg-gray-50 p-4 rounded-lg"
          />
        )
      
      case 'code':
        return (
          <CodeBlock
            code={section.content}
            language={section.language || 'javascript'}
            title={section.title}
          />
        )
      
      case 'image':
        return (
          <div className="my-6">
            <Image
              src={section.content}
              alt={section.title || 'Educational content'}
              width={600}
              height={400}
              className="rounded-lg shadow-sm"
            />
            {section.title && (
              <p className="text-sm text-gray-600 mt-2 text-center italic">
                {section.title}
              </p>
            )}
          </div>
        )
      
      case 'diagram':
        return (
          <div className="my-6 p-6 bg-gray-50 rounded-lg">
            <div className="flex items-center justify-center h-48 bg-white rounded border-2 border-dashed border-gray-300">
              <div className="text-center">
                <p className="text-gray-500 mb-2">Interactive Diagram</p>
                <p className="text-sm text-gray-400">{section.content}</p>
              </div>
            </div>
            {section.title && (
              <p className="text-sm text-gray-600 mt-2 text-center">
                {section.title}
              </p>
            )}
          </div>
        )
      
      default:
        return <p className="text-gray-700">{section.content}</p>
    }
  }

  return (
    <div className="mb-6">
      {section.title && (
        <h3 className="text-lg font-semibold text-gray-900 mb-3">
          {section.title}
        </h3>
      )}
      {renderContent()}
    </div>
  )
}