# Define naming moniker
locals {
  moniker = "sonic-sync"
}

# Create VPC
resource "aws_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support = true
  tags = {
    Name = "${local.moniker}-vpc"
  }
}

# Create public subnets
resource "aws_subnet" "public_subnet_1" {
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-east-1a"
  map_public_ip_on_launch = true
  ipv6_cidr_block_association {
    ipv6_cidr_block = "2001:db8:1::/64"
  }
  tags = {
    Name = "${local.moniker}-public-subnet-1"
  }
}

resource "aws_subnet" "public_subnet_2" {
  vpc_id = aws_vpc.vpc.id
  cidr_block = "10.0.2.0/24"
  availability_zone = "us-east-1b"
  map_public_ip_on_launch = true
  ipv6_cidr_block_association {
    ipv6_cidr_block = "2001:db8:2::/64"
  }
  tags = {
    Name = "${local.moniker}-public-subnet-2"
  }
}

# Create security group for ALB
resource "aws_security_group" "alb_sg" {
  vpc_id    = aws_vpc.vpc.id
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow HTTP traffic"
  }
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow HTTPS traffic"
  }
  tags = {
    Name = "${local.moniker}-alb-security-group"
  }
}

# Create ALB
resource "aws_lb" "alb" {
  name               = "${local.moniker}-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb_sg.id]
  subnets            = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  tags = {
    Name = "${local.moniker}-alb"
  }
}

# Create target group
resource "aws_lb_target_group" "tg" {
  name        = "${local.moniker}-target-group"
  port        = 80
  protocol    = "HTTP"
  vpc_id      = aws_vpc.vpc.id
  target_type = "ip"
  tags = {
    Name = "${local.moniker}-target-group"
  }
}

# Create ECS Fargate task definition
resource "aws_ecs_task_definition" "task" {
  family                   = "${local.moniker}-task"
  execution_role_arn       = aws_iam_role.task_execution_role.arn
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
  requires_compatibilities = ["FARGATE"]

  container_definitions = <<CONTAINER_DEF
[
  {
    "name": "${local.moniker}-container",
    "image": "public_ecr_image_url_here",
    "portMappings": [
      {
        "containerPort": 80,
        "protocol": "tcp"
      },
      {
        "containerPort": 443,
        "protocol": "tcp"
      }
    ],
    "logConfiguration": {
      "logDriver": "awslogs",
      "options": {
        "awslogs-group": "/ecs/${local.moniker}-task",
        "awslogs-region": "us-east-1",
        "awslogs-stream-prefix": "${local.moniker}-container"
      }
    }
  }
]
CONTAINER_DEF

  tags = {
    Name = "${local.moniker}-task"
  }
}

# Create ECS cluster
resource "aws_ecs_cluster" "cluster" {
  name = "${local.moniker}-cluster"

  tags = {
    Name = "${local.moniker}-cluster"
  }
}

# Create ECS service
resource "aws_ecs_service" "service" {
  name            = "${local.moniker}-service"
  cluster         = aws_ecs_cluster.cluster.id
  task_definition = aws_ecs_task_definition.task.arn
  desired_count   = 1
}

# Create ALB listener rule to associate with the target group
resource "aws_lb_listener_rule" "rule" {
  listener_arn = aws_lb_listener.http.arn
  priority     = 1
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.tg.arn
  }
}
