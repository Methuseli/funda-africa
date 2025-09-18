export interface Subject {
  id: string
  title: string
  description: string
  icon: string
  color: string
  topics: Topic[]
}

export interface Topic {
  id: string
  title: string
  description: string
  subtopics: Subtopic[]
}

export interface Subtopic {
  id: string
  title: string
  description: string
  content: SubtopicContent
  estimatedTime: number
  difficulty: 'Beginner' | 'Intermediate' | 'Advanced'
}

export interface SubtopicContent {
  sections: ContentSection[]
  aiGeneratedNotes?: string
  aiGeneratedGraphs?: string[]
  exercises?: Exercise[]
}

export interface ContentSection {
  id: string
  type: 'text' | 'latex' | 'code' | 'image' | 'diagram'
  title?: string
  content: string
  language?: string
}

export interface Exercise {
  id: string
  question: string
  options?: string[]
  correctAnswer: string | number
  explanation: string
}

export const subjects: Subject[] = [
  {
    id: 'mathematics',
    title: 'Mathematics',
    description: 'From basic algebra to advanced calculus',
    icon: 'Calculator',
    color: 'bg-blue-500',
    topics: [
      {
        id: 'linear-algebra',
        title: 'Linear Algebra',
        description: 'Vectors, matrices, and linear transformations',
        subtopics: [
          {
            id: 'vectors-basics',
            title: 'Vector Basics',
            description: 'Introduction to vectors and vector operations',
            estimatedTime: 45,
            difficulty: 'Beginner',
            content: {
              sections: [
                {
                  id: '1',
                  type: 'text',
                  title: 'What is a Vector?',
                  content: 'A vector is a mathematical object that has both magnitude and direction. In physics and engineering, vectors are used to represent quantities like velocity, force, and displacement.'
                },
                {
                  id: '2',
                  type: 'latex',
                  title: 'Vector Notation',
                  content: 'A vector $\\vec{v}$ in 2D space can be written as $\\vec{v} = \\langle v_x, v_y \\rangle$ or in component form as $\\vec{v} = v_x\\hat{i} + v_y\\hat{j}$'
                },
                {
                  id: '3',
                  type: 'text',
                  title: 'Vector Addition',
                  content: 'Vectors can be added component-wise. If we have two vectors $\\vec{a} = \\langle a_x, a_y \\rangle$ and $\\vec{b} = \\langle b_x, b_y \\rangle$, then their sum is:'
                },
                {
                  id: '4',
                  type: 'latex',
                  content: '$\\vec{a} + \\vec{b} = \\langle a_x + b_x, a_y + b_y \\rangle$'
                },
                {
                  id: '5',
                  type: 'code',
                  title: 'Vector Operations in Python',
                  content: 'import numpy as np\n\n# Define two vectors\nvector_a = np.array([3, 4])\nvector_b = np.array([1, 2])\n\n# Vector addition\nsum_vector = vector_a + vector_b\nprint(f"Sum: {sum_vector}")\n\n# Vector magnitude\nmagnitude_a = np.linalg.norm(vector_a)\nprint(f"Magnitude of a: {magnitude_a}")',
                  language: 'python'
                }
              ],
              aiGeneratedNotes: 'Key Points to Remember:\n• Vectors have both magnitude and direction\n• Vector addition follows the parallelogram law\n• Unit vectors have magnitude 1\n• The zero vector has magnitude 0 and no defined direction',
              exercises: [
                {
                  id: '1',
                  question: 'What is the magnitude of vector $\\vec{v} = \\langle 3, 4 \\rangle$?',
                  options: ['5', '7', '12', '25'],
                  correctAnswer: 0,
                  explanation: 'The magnitude is calculated using $|\\vec{v}| = \\sqrt{3^2 + 4^2} = \\sqrt{9 + 16} = \\sqrt{25} = 5$'
                }
              ]
            }
          }
        ]
      },
      {
        id: 'calculus',
        title: 'Calculus',
        description: 'Limits, derivatives, and integrals',
        subtopics: [
          {
            id: 'limits-intro',
            title: 'Introduction to Limits',
            description: 'Understanding the concept of limits in calculus',
            estimatedTime: 60,
            difficulty: 'Intermediate',
            content: {
              sections: [
                {
                  id: '1',
                  type: 'text',
                  title: 'What is a Limit?',
                  content: 'In calculus, a limit describes the behavior of a function as its input approaches a particular value. Limits are fundamental to defining derivatives and integrals.'
                },
                {
                  id: '2',
                  type: 'latex',
                  title: 'Limit Notation',
                  content: 'We write $\\lim_{x \\to a} f(x) = L$ to mean "the limit of f(x) as x approaches a equals L"'
                }
              ],
              aiGeneratedNotes: 'Essential Limit Concepts:\n• Limits describe function behavior near a point\n• One-sided limits may differ\n• Limits may exist even when the function is undefined at that point\n• The squeeze theorem is useful for difficult limits'
            }
          }
        ]
      }
    ]
  },
  {
    id: 'physics',
    title: 'Physics',
    description: 'Understanding the natural world through scientific principles',
    icon: 'Zap',
    color: 'bg-green-500',
    topics: [
      {
        id: 'mechanics',
        title: 'Classical Mechanics',
        description: 'Motion, forces, and energy',
        subtopics: [
          {
            id: 'newtons-laws',
            title: "Newton's Laws of Motion",
            description: 'The three fundamental laws governing motion',
            estimatedTime: 90,
            difficulty: 'Beginner',
            content: {
              sections: [
                {
                  id: '1',
                  type: 'text',
                  title: "Newton's First Law (Law of Inertia)",
                  content: 'An object at rest stays at rest, and an object in motion stays in motion at constant velocity, unless acted upon by an external force.'
                },
                {
                  id: '2',
                  type: 'latex',
                  title: "Newton's Second Law",
                  content: '$\\vec{F} = m\\vec{a}$ where $\\vec{F}$ is the net force, $m$ is mass, and $\\vec{a}$ is acceleration.'
                }
              ],
              aiGeneratedNotes: 'Newton\'s Laws Summary:\n• First Law: Objects resist changes in motion (inertia)\n• Second Law: Force equals mass times acceleration\n• Third Law: Every action has an equal and opposite reaction\n• These laws only apply in inertial reference frames'
            }
          }
        ]
      }
    ]
  },
  {
    id: 'computer-science',
    title: 'Computer Science',
    description: 'Programming, algorithms, and computational thinking',
    icon: 'Code',
    color: 'bg-purple-500',
    topics: [
      {
        id: 'algorithms',
        title: 'Algorithms & Data Structures',
        description: 'Efficient problem-solving techniques',
        subtopics: [
          {
            id: 'sorting-algorithms',
            title: 'Sorting Algorithms',
            description: 'Various methods to arrange data in order',
            estimatedTime: 120,
            difficulty: 'Intermediate',
            content: {
              sections: [
                {
                  id: '1',
                  type: 'text',
                  title: 'Introduction to Sorting',
                  content: 'Sorting algorithms arrange elements in a specific order (usually ascending or descending). Different algorithms have different time and space complexities.'
                },
                {
                  id: '2',
                  type: 'code',
                  title: 'Bubble Sort Implementation',
                  content: 'def bubble_sort(arr):\n    n = len(arr)\n    for i in range(n):\n        for j in range(0, n - i - 1):\n            if arr[j] > arr[j + 1]:\n                arr[j], arr[j + 1] = arr[j + 1], arr[j]\n    return arr\n\n# Example usage\nnumbers = [64, 34, 25, 12, 22, 11, 90]\nsorted_numbers = bubble_sort(numbers.copy())\nprint(f"Sorted array: {sorted_numbers}")',
                  language: 'python'
                }
              ],
              aiGeneratedNotes: 'Sorting Algorithm Comparison:\n• Bubble Sort: O(n²) time, simple but inefficient\n• Quick Sort: O(n log n) average, O(n²) worst case\n• Merge Sort: O(n log n) guaranteed, stable\n• Selection Sort: O(n²) time, minimal swaps'
            }
          }
        ]
      }
    ]
  }
]

export const getSubject = (id: string) => subjects.find(s => s.id === id)
export const getTopic = (subjectId: string, topicId: string) => 
  getSubject(subjectId)?.topics.find(t => t.id === topicId)
export const getSubtopic = (subjectId: string, topicId: string, subtopicId: string) =>
  getTopic(subjectId, topicId)?.subtopics.find(s => s.id === subtopicId)