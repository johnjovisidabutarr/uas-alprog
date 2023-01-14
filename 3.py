from typing import List

class Student:
    def __init__(self, name: str, mtk: int, bahasa: int, ipa: int):
        self.name = name
        self.mtk = mtk
        self.bahasa = bahasa
        self.ipa = ipa
    
    def average(self) -> float:
        return (self.mtk + self.bahasa + self.ipa) / 3

class Classroom:
    def __init__(self, students: List[Student]):
        self.students = students
    
    def top_3(self) -> List[Student]:
        self.students.sort(key=lambda x: x.average(), reverse=True)
        return self.students[:3]
    
    def find_average(self, name: str) -> float:
        for student in self.students:
            if student.name == name:
                return student.average()
        return -1

def main():
    students = []
    name = ""
    mtk = 0
    bahasa = 0
    ipa = 0
    
    while True:
        name, mtk, bahasa, ipa = input().split()
        if name == "stop":
            break
        students.append(Student(name, int(mtk), int(bahasa), int(ipa)))
    
    classroom = Classroom(students)
    student_name = input()
    average = classroom.find_average(student_name)
    if average != -1:
        print("Nilai rata-rata dari", student_name, "adalah", average)
    top_3_students = classroom.top_3()
    for i, student in enumerate(top_3_students):
        print(student.name, "rangking",i+1, "dengan rata-rata", student.average())

if __name__ == "__main__":
    main()
