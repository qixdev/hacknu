import os
from flask import Flask, request, jsonify
import matplotlib.pyplot as plt
import matplotlib


matplotlib.use('Agg')

app = Flask(__name__)

if not os.path.exists('static'):
    os.makedirs('static')


@app.route('/graphics/', methods=['POST'])
def graphics():
    data = request.get_json()
    name = data['name']
    grades = []
    subjects = []
    files = []

    for subject_name, subject_data in data['materials'].items():
        topics = subject_data['topics']
        grade = subject_data['grade']
        grades.append(grade)
        subjects.append(subject_name)

        topic_names = list(topics.keys())
        topic_grades = [topic['averageGrade'] for topic in topics.values()]

        plt.figure(figsize=(10, 6))
        plt.plot(topic_names, topic_grades, marker='o',
                 linestyle='-', color='skyblue')
        plt.xlabel('Topics')
        plt.ylabel('Average Grades')
        plt.title(f'Average Grades for {subject_name} by Topics')
        plt.xticks(rotation=90)
        plt.ylim(0, 100)
        plt.yticks(range(0, 101, 10))
        plt.tight_layout()

        filename = f'static/{subject_name}_grades_{name}.png'
        plt.savefig(filename)
        plt.close()
        files.append(f'/static/{subject_name}_grades_{name}.png')

    plt.figure(figsize=(10, 6))
    plt.bar(subjects, grades, color='skyblue')
    plt.xlabel('Subjects')
    plt.ylabel('Scores')
    plt.title('Student Subject Scores')
    plt.xticks(rotation=90)
    plt.ylim(0, 100)
    plt.yticks(range(0, 101, 10))
    plt.tight_layout()
    filename = f'static/subjects_grades_{name}.png'
    plt.savefig(filename)
    plt.close()
    files.append(f'/static/subjects_grades_{name}.png')

    return jsonify({"message": "Graphs generated successfully", "files": files, "name": name})


if __name__ == '__main__':
    app.run(debug=True, port=3000)