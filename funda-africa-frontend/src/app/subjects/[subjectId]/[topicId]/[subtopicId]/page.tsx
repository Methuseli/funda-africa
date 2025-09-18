"use client"

import React, { useEffect } from 'react'
import Link from 'next/link'
import { useParams } from 'next/navigation'
import { ArrowLeft, ArrowRight, Clock, Award, BookOpen, Lightbulb } from 'lucide-react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Separator } from '@/components/ui/separator'
import { Alert, AlertDescription } from '@/components/ui/alert'
import Sidebar from '@/components/layout/sidebar'
import ContentSectionComponent from '@/components/content/content-section'
import { getSubtopic } from '@/lib/data'
import { useAppStore } from '@/lib/store'

export default function SubtopicPage() {
  const params = useParams(), target: 5 },
    { goal: 'S
  const subjectId = params.subjectId as string
  const topicId = params.topicId as string
  const subtopicId = params.subtopicId as string
  
  const subtopic = getSubtopic(subjectId, topicId, subtopicId)
  const { setCurrentSubtopic } = useAppStore()

  useEffect(() => {
    setCurrentSubtopic(subtopicId)
    return () => setCurrentSubtopic(null)
  }, [subtopicId, setCurrentSubtopic])

  if (!subtopic) {
    return (
      <div className="flex w-full">
        <Sidebar />
        <div className="flex-1 p-8">
          <div className="text-center">
            <h1 className="text-2xl font-bold text-gray-900">Subtopic not found</h1>
            <p className="text-gray-600 mt-2">The subtopic you're looking for doesn't exist.</p>
            <Link href={`/subjects/${subjectId}/${topicId}`}>
              <Button className="mt-4">Back to Topic</Button>
            </Link>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="flex w-full">
      <Sidebar />
      
      <div className="flex-1">
        <div className="max-w-4xl mx-auto p-8 space-y-8">
          {/* Breadcrumb */}
          <div className="flex items-center space-x-2 text-sm text-gray-600">
            <Link href="/" className="hover:text-gray-900">Dashboard</Link>
            <span>/</span>
            <Link href={`/subjects/${subjectId}`} className="hover:text-gray-900">
              {subjectId.charAt(0).toUpperCase() + subjectId.slice(1).replace('-', ' ')}
            </Link>
            <span>/</span>
            <Link href={`/subjects/${subjectId}/${topicId}`} className="hover:text-gray-900">
              {topicId.charAt(0).toUpperCase() + topicId.slice(1).replace('-', ' ')}
            </Link>
            <span>/</span>
            <span className="text-gray-900">{subtopic.title}</span>
          </div>

          {/* Navigation */}
          <div className="flex items-center justify-between">
            <Link href={`/subjects/${subjectId}/${topicId}`}>
              <Button variant="ghost">
                <ArrowLeft className="h-4 w-4 mr-2" />
                Back to Topic
              </Button>
            </Link>
            
            <div className="flex items-center space-x-2">
              <Badge 
                variant={
                  subtopic.difficulty === 'Beginner' ? 'secondary' :
                  subtopic.difficulty === 'Intermediate' ? 'default' : 'destructive'
                }
              >
                {subtopic.difficulty}
              </Badge>
              <div className="flex items-center space-x-1 text-sm text-gray-600">
                <Clock className="h-4 w-4" />
                <span>{subtopic.estimatedTime} min</span>
              </div>
            </div>
          </div>

          {/* Header */}
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold text-gray-900">{subtopic.title}</h1>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">{subtopic.description}</p>
          </div>

          {/* Main Content */}
          <div className="space-y-8">
            {/* Learning Content */}
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center space-x-2">
                  <BookOpen className="h-5 w-5" />
                  <span>Lesson Content</span>
                </CardTitle>
              </CardHeader>
              <CardContent className="prose prose-gray max-w-none">
                <div className="space-y-6">
                  {subtopic.content.sections.map((section) => (
                    <ContentSectionComponent key={section.id} section={section} />
                  ))}
                </div>
              </CardContent>
            </Card>

            {/* AI Generated Notes */}
            {subtopic.content.aiGeneratedNotes && (
              <Alert className="bg-blue-50 border-blue-200">
                <Lightbulb className="h-4 w-4 text-blue-600" />
                <AlertDescription className="text-blue-800">
                  <div className="space-y-2">
                    <h4 className="font-semibold">AI-Generated Study Notes</h4>
                    <div className="whitespace-pre-line text-sm">{subtopic.content.aiGeneratedNotes}</div>
                  </div>
                </AlertDescription>
              </Alert>
            )}

            {/* Practice Exercises */}
            {subtopic.content.exercises && subtopic.content.exercises.length > 0 && (
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center space-x-2">
                    <Award className="h-5 w-5" />
                    <span>Practice Exercises</span>
                  </CardTitle>
                  <CardDescription>
                    Test your understanding with these practice questions
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <div className="space-y-6">
                    {subtopic.content.exercises.map((exercise, index) => (
                      <div key={exercise.id} className="p-6 bg-gray-50 rounded-lg">
                        <h4 className="font-semibold mb-3">Question {index + 1}</h4>
                        <p className="text-gray-700 mb-4">{exercise.question}</p>
                        
                        {exercise.options && (
                          <div className="space-y-2 mb-4">
                            {exercise.options.map((option, optionIndex) => (
                              <div key={optionIndex} className="flex items-center space-x-2">
                                <div className="w-6 h-6 rounded-full border-2 border-gray-300 flex items-center justify-center text-sm font-medium">
                                  {String.fromCharCode(65 + optionIndex)}
                                </div>
                                <span>{option}</span>
                              </div>
                            ))}
                          </div>
                        )}
                        
                        <details className="mt-4">
                          <summary className="cursor-pointer text-blue-600 hover:text-blue-800 font-medium">
                            Show Answer & Explanation
                          </summary>
                          <div className="mt-3 p-4 bg-white rounded border-l-4 border-green-500">
                            <p className="text-green-700 font-medium mb-2">
                              Answer: {exercise.options ? 
                                `${String.fromCharCode(65 + Number(exercise.correctAnswer))} - ${exercise.options[Number(exercise.correctAnswer)]}` : 
                                exercise.correctAnswer}
                            </p>
                            <p className="text-gray-700">{exercise.explanation}</p>
                          </div>
                        </details>
                      </div>
                    ))}
                  </div>
                </CardContent>
              </Card>
            )}

            {/* Navigation Footer */}
            <div className="flex items-center justify-between pt-8 border-t">
              <Button variant="outline">
                <ArrowLeft className="h-4 w-4 mr-2" />
                Previous Subtopic
              </Button>
              
              <div className="text-center">
                <p className="text-sm text-gray-600 mb-2">Mark as complete to continue</p>
                <Button size="lg" className="bg-green-600 hover:bg-green-700">
                  Mark Complete & Continue
                  <ArrowRight className="h-4 w-4 ml-2" />
                </Button>
              </div>
              
              <Button variant="outline">
                Next Subtopic
                <ArrowRight className="h-4 w-4 ml-2" />
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}