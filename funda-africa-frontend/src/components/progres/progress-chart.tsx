"use client"

import React from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Progress } from '@/components/ui/progress'
import { 
  ResponsiveContainer, 
  LineChart, 
  Line, 
  XAxis, 
  YAxis, 
  CartesianGrid, 
  Tooltip, 
  PieChart, 
  Pie, 
  Cell,
  BarChart,
  Bar,
  Legend
} from 'recharts'

const progressData = [
  { day: 'Mon', progress: 20 },
  { day: 'Tue', progress: 35 },
  { day: 'Wed', progress: 42 },
  { day: 'Thu', progress: 58 },
  { day: 'Fri', progress: 65 },
  { day: 'Sat', progress: 78 },
  { day: 'Sun', progress: 85 },
]

const subjectProgress = [
  { name: 'Mathematics', value: 75, color: '#3B82F6' },
  { name: 'Physics', value: 60, color: '#10B981' },
  { name: 'Computer Science', value: 85, color: '#8B5CF6' },
]

const weeklyActivity = [
  { subject: 'Math', hoursSpent: 12 },
  { subject: 'Physics', hoursSpent: 8 },
  { subject: 'CS', hoursSpent: 15 },
]

export default function ProgressChart() {
  const totalProgress = Math.round(
    subjectProgress.reduce((acc, subject) => acc + subject.value, 0) / subjectProgress.length
  )

  return (
    <div className="space-y-6">
      {/* Overall Progress */}
      <Card>
        <CardHeader>
          <CardTitle>Overall Progress</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <span className="text-sm font-medium">Total Completion</span>
              <span className="text-2xl font-bold text-blue-600">{totalProgress}%</span>
            </div>
            <Progress value={totalProgress} className="h-3" />
            
            <div className="space-y-3 mt-6">
              {subjectProgress.map((subject) => (
                <div key={subject.name} className="space-y-2">
                  <div className="flex items-center justify-between">
                    <span className="text-sm font-medium">{subject.name}</span>
                    <span className="text-sm text-gray-600">{subject.value}%</span>
                  </div>
                  <Progress value={subject.value} className="h-2" />
                </div>
              ))}
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Weekly Progress Trend */}
      <Card>
        <CardHeader>
          <CardTitle>Weekly Progress Trend</CardTitle>
        </CardHeader>
        <CardContent>
          <ResponsiveContainer width="100%" height={200}>
            <LineChart data={progressData}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="day" />
              <YAxis />
              <Tooltip />
              <Line 
                type="monotone" 
                dataKey="progress" 
                stroke="#3B82F6" 
                strokeWidth={3}
                dot={{ fill: '#3B82F6' }}
              />
            </LineChart>
          </ResponsiveContainer>
        </CardContent>
      </Card>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {/* Subject Distribution */}
        <Card>
          <CardHeader>
            <CardTitle>Subject Progress Distribution</CardTitle>
          </CardHeader>
          <CardContent>
            <ResponsiveContainer width="100%" height={200}>
              <PieChart>
                <Pie
                  data={subjectProgress}
                  cx="50%"
                  cy="50%"
                  outerRadius={60}
                  fill="#8884d8"
                  dataKey="value"
                  label={({ name, value }) => `${name}: ${value}%`}
                  labelLine={false}
                >
                  {subjectProgress.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.color} />
                  ))}
                </Pie>
                <Tooltip />
              </PieChart>
            </ResponsiveContainer>
          </CardContent>
        </Card>

        {/* Weekly Activity */}
        <Card>
          <CardHeader>
            <CardTitle>Weekly Study Hours</CardTitle>
          </CardHeader>
          <CardContent>
            <ResponsiveContainer width="100%" height={200}>
              <BarChart data={weeklyActivity}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="subject" />
                <YAxis />
                <Tooltip />
                <Bar dataKey="hoursSpent" fill="#3B82F6" radius={[4, 4, 0, 0]} />
              </BarChart>
            </ResponsiveContainer>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}